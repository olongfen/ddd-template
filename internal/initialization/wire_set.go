package initialization

import (
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/xlog"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var Set = wire.NewSet(InitLog, InitConf)

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
