package xlog

import (
	"ddd-template/common/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func encodeJSON() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func writer() zapcore.WriteSyncer {
	logCfg := conf.Get().Log
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logCfg.Filename,
		MaxSize:    logCfg.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: logCfg.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     logCfg.MaxAges,    // 保留旧文件的最大天数
		Compress:   logCfg.Compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func NewProduceLogger() *zap.Logger {
	core := zapcore.NewCore(encodeJSON(), writer(), zapcore.DebugLevel)
	return zap.New(core, zap.AddCaller())
}

func NewDevelopment() *zap.Logger {
	log, _ := zap.NewProduction()
	return log
}
