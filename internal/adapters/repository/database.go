package repository

import (
	"context"
	"ddd-template/internal/adapters/repository/db_iface"
	"ddd-template/internal/rely"
	"fmt"
	"github.com/olongfen/toolkit/err_mul"
	"github.com/olongfen/toolkit/scontext"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

// GormSpanKey 包内静态变量
const GormSpanKey = "__gorm_span"
const (
	tableNamePrefix        = ""
	RepositoryMethodCtxTag = "repository_method"
	CallBackBeforeName     = "opentracing:before"
	CallBackAfterName      = "opentracing:after"
)

type Data struct {
	db  *gorm.DB
	log *zap.Logger
}

type contextTxKey struct{}

func NewTransaction(d db_iface.DBData) db_iface.ITransaction {
	return d
}

func (d *Data) ExecTx(ctx context.Context, fc func(context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fc(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db.WithContext(ctx)
}

func (d *Data) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func NewData(db *gorm.DB, logger *zap.Logger) (ret db_iface.DBData) {
	return &Data{
		db:  db,
		log: logger,
	}
}

// InitDBConnect init database connect
func InitDBConnect(c *rely.Configs, logger *zap.Logger) (res *gorm.DB) {
	var (
		err        error
		db         *gorm.DB
		gormConfig = &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: tableNamePrefix,
			},
			//Logger: utils.New(lg,utils.Config{Colorful: true}),
		}
	)
	switch c.Database.Driver {
	case "postgresql":
		dbCfg := c.Database
		dns := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`, dbCfg.Host,
			dbCfg.User, dbCfg.Password, dbCfg.DBName, dbCfg.Port, dbCfg.SSLMode, dbCfg.TimeZone)
		if db, err = gorm.Open(postgres.Open(dns), gormConfig); err != nil {
			logger.Fatal(err.Error())
		}

	}

	if db == nil {
		logger.Sugar().Fatal("data do not init")
		return
	}
	err = db.Use(&OpentracingPlugin{})
	if err != nil {
		logger.Fatal(err.Error())

	}
	// true 自动迁移
	if c.Database.AutoMigrate {
		err = db.AutoMigrate()
		if err != nil {
			logger.Fatal(err.Error())
		}
	}
	// debug 打开数据库debug打印模式
	if rely.Get().Database.Debug {
		db = db.Debug()
	}

	return db
}

type OpentracingPlugin struct {
}

var _ gorm.Plugin = &OpentracingPlugin{}

func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前 - 并不是都用相同的方法，可以自己自定义
	if err = db.Callback().Create().Before("gorm:before_create").Register(CallBackBeforeName, before); err != nil {
		return
	}
	if err = db.Callback().Query().Before("gorm:query").Register(CallBackBeforeName, before); err != nil {
		return
	}
	if err = db.Callback().Delete().Before("gorm:before_delete").Register(CallBackBeforeName, before); err != nil {
		return
	}
	if err = db.Callback().Update().Before("gorm:setup_reflect_value").Register(CallBackBeforeName, before); err != nil {
		return
	}
	if err = db.Callback().Row().Before("gorm:row").Register(CallBackBeforeName, before); err != nil {
		return
	}
	if err = db.Callback().Raw().Before("gorm:raw").Register(CallBackBeforeName, before); err != nil {
		return
	}

	// 结束后 - 并不是都用相同的方法，可以自己自定义
	if err = db.Callback().Create().After("gorm:after_create").Register(CallBackAfterName, after); err != nil {
		return
	}
	if err = db.Callback().Query().After("gorm:after_query").Register(CallBackAfterName, after); err != nil {
		return
	}
	if err = db.Callback().Delete().After("gorm:after_delete").Register(CallBackAfterName, after); err != nil {
		return
	}
	if err = db.Callback().Update().After("gorm:after_update").Register(CallBackAfterName, after); err != nil {
		return
	}
	if err = db.Callback().Row().After("gorm:row").Register(CallBackAfterName, after); err != nil {
		return
	}
	if err = db.Callback().Raw().After("gorm:raw").Register(CallBackAfterName, after); err != nil {
		return
	}
	return
}

func before(db *gorm.DB) {

	tr := otel.Tracer("gorm-before")

	_, span := tr.Start(db.Statement.Context, "gorm-before")
	// 利用db实例去传递span
	db.InstanceSet(GormSpanKey, span)

}

func after(db *gorm.DB) {
	if db.Error != nil {
		handlerDBError(db)
	}
	_span, exist := db.InstanceGet(GormSpanKey)
	if !exist {
		return
	}
	// 断言类型
	span, ok := _span.(trace.Span)
	if !ok {
		return
	}

	defer span.End()

	if db.Error != nil {
		span.SetAttributes(attribute.Key("gorm_err").String(db.Error.Error()))
	}

	span.SetAttributes(attribute.Key("sql").String(db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))

}

func handlerDBError(db *gorm.DB) {
	lang := scontext.GetLanguage(db.Statement.Context)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		db.Error = err_mul.NewError(err_mul.RecordNotFound, lang)
		return
	}
	msg := db.Error.Error()
	const (
		code23505 = "23505"
	)
	// 处理数据库错误
	for _, v := range db.Statement.Schema.DBNames {
		if strings.Contains(msg, code23505) && strings.Contains(msg, v) {
			field := db.Statement.Schema.FieldsByDBName[v]
			name := strings.ToLower(field.Name[:1]) + field.Name[1:]
			var errs = err_mul.DBErrorResponse{}
			errs[name] = err_mul.NewError(err_mul.AlreadyExists, lang)
			db.Error = errs
		}
	}
}

/**
 * 驼峰转蛇形 snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 * @date 2020/7/30
 * @param s 需要转换的字符串
 * @return string
 **/
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
