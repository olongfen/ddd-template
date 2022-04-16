package initialization

import (
	"ddd-template/internal/common/xlog"
	"ddd-template/internal/initialization/conf"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var Set = wire.NewSet(InitLog, conf.InitConf)

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
