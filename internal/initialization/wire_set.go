package initialization

import (
	"ddd-template/internal/common/xlog"
	"ddd-template/internal/initialization/conf"
	"ddd-template/internal/initialization/database"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var Set = wire.NewSet(database.NewDatabase, InitLog, conf.InitConf)

func InitLog(cfg *conf.Configs) *zap.Logger {
	var (
		logger *zap.Logger
	)
	if cfg.Environment == "dev" {
		logger = xlog.NewDevelopment()
	} else {
		logger = xlog.NewProduceLogger()
	}
	xlog.Log = logger
	return logger
}
