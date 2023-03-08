package rely

import (
	"github.com/olongfen/toolkit/xlog"
	"go.uber.org/zap"
)

var (
	GlobalLogger = xlog.NewDevelopment()
)

//
//const (
//	ProjectNotBelongToPlatform = 43001
//	GenPDFFailed               = 43002
//)

//func init() {
//	xerror.DefaultErrorMul.Set(ProjectNotBelongToPlatform, "zh", "项目不属于该平台")
//	xerror.DefaultErrorMul.Set(ProjectNotBelongToPlatform, "en", "item does not belong to this platform")
//	xerror.DefaultErrorMul.Set(GenPDFFailed, "zh", "生成pdf失败")
//	xerror.DefaultErrorMul.Set(GenPDFFailed, "en", "gen pdf failed")
//}

// NewLogger new log
func NewLogger(cfg *Configs) *zap.Logger {
	if !cfg.Log.Debug {
		GlobalLogger = xlog.NewProduceLogger(xlog.Config{
			InfoFile:   cfg.Log.Filename,
			ErrorFile:  cfg.Log.ErrorFile,
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAges,
			Compress:   cfg.Log.Compress,
		})

	}
	return GlobalLogger
}
