package repository

import (
	"context"
	"ddd-template/internal/app"
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/xlog"
	"ddd-template/internal/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func NewTransaction(d *Data) app.ITransaction {
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
	return d.db
}

func NewData(db *gorm.DB, logger *zap.Logger) (ret *Data) {
	return &Data{
		db:  db,
		log: logger,
	}
}

// NewDB new
func NewDB(c *conf.Configs, logger *zap.Logger) (res *gorm.DB) {
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
	case "postgres":
		if db, err = gorm.Open(postgres.Open(c.Database.Source), gormConfig); err != nil {
			logger.Sugar().Fatal(err.Error())
		}

	}

	if db == nil {
		logger.Sugar().Fatal("data do not init")
		return
	}
	err = db.Use(&OpentracingPlugin{})
	if err != nil {
		logger.Sugar().Fatal(err)
		return
	}
	if conf.Get().Database.Dev {
		db = db.Debug()
		err = db.AutoMigrate(&domain.Demo{})
		if err != nil {
			xlog.Log.Sugar().Warn(err)
			err = nil
		}
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
	db.Callback().Create().Before("gorm:before_create").Register(CallBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(CallBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(CallBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(CallBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(CallBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(CallBackBeforeName, before)

	// 结束后 - 并不是都用相同的方法，可以自己自定义
	db.Callback().Create().After("gorm:after_create").Register(CallBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(CallBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(CallBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(CallBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(CallBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(CallBackAfterName, after)
	return
}

func before(db *gorm.DB) {

	tr := otel.Tracer("gorm-before")

	_, span := tr.Start(db.Statement.Context, "gorm-before")
	// 利用db实例去传递span
	db.InstanceSet(GormSpanKey, span)

}

func after(db *gorm.DB) {
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
