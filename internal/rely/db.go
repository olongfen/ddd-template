package rely

import (
	"ddd-template/internal/domain"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/olongfen/toolkit/db_data"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	// tableNamePrefix 表格头部，需要自己定义
	tableNamePrefix = ""
)

// InitDBConnect init database connect
func InitDBConnect(c *Configs, logger *zap.Logger) (res *gorm.DB, err error) {
	var (
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
			return
		}

	}

	if db == nil {
		err = errors.New("db dose not init")
		return
	}
	// true 自动迁移
	if c.Database.AutoMigrate {
		err = db.AutoMigrate(&domain.Demo{})
		if v, ok := err.(*pgconn.PgError); ok {
			if v.Code == "42P07" {
				err = nil
			}
		}
		if err != nil {
			return
		}
	}
	err = db.Use(&db_data.OpentracingPlugin{})
	if err != nil {
		return

	}
	// debug 打开数据库debug打印模式
	if c.Database.Debug {
		db = db.Debug()
	}

	return db, nil
}
