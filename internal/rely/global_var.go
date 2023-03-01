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
//	err_mul.DefaultErrorMul.Set(ProjectNotBelongToPlatform, "zh", "项目不属于该平台")
//	err_mul.DefaultErrorMul.Set(ProjectNotBelongToPlatform, "en", "item does not belong to this platform")
//	err_mul.DefaultErrorMul.Set(GenPDFFailed, "zh", "生成pdf失败")
//	err_mul.DefaultErrorMul.Set(GenPDFFailed, "en", "gen pdf failed")
//}

// NewLogger new log
func NewLogger(cfg *Configs) *zap.Logger {
	if !cfg.Log.Debug {
		GlobalLogger = xlog.NewProduceLogger()

	}
	return GlobalLogger
}
