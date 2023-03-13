package rely

import (
	"fmt"
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
func InitDBConnect(c *Configs, logger *zap.Logger) (res *gorm.DB) {
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
	err = db.Use(&db_data.OpentracingPlugin{})
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
	if c.Database.Debug {
		db = db.Debug()
	}

	return db
}
