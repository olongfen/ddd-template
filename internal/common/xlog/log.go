package xlog

import (
	"context"
	"ddd-template/internal/initialization/conf"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

var Log = NewDevelopment()

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

type DBLog struct {
	*zap.Logger
	LogLevel                            zapcore.Level
	SlowThreshold                       time.Duration
	IgnoreRecordNotFoundError           bool
	traceStr, traceErrStr, traceWarnStr string
}

func (l *DBLog) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = zapcore.Level(level)
	return &newlogger
}

func (l *DBLog) Info(ctx context.Context, s string, i ...interface{}) {
	l.Sugar().Info(s, i)
}

func (l *DBLog) Warn(ctx context.Context, s string, i ...interface{}) {
	l.Sugar().Warn(s, i)
}

func (l *DBLog) Error(ctx context.Context, s string, i ...interface{}) {
	l.Sugar().Error()
}

func (l *DBLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= zapcore.DebugLevel {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= zapcore.ErrorLevel && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Sugar().Infof(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Sugar().Infof(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= zap.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Sugar().Infof(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Sugar().Infof(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == zap.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Sugar().Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Sugar().Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func NewDBLog(zapLog *zap.Logger) logger.Interface {
	var (
		traceStr     = `%s [%.3fms] [rows:%v] %s`
		traceWarnStr = `%s %s[%.3fms] [rows:%v] %s`
		traceErrStr  = `%s %s[%.3fms] [rows:%v] %s`
	)
	return &DBLog{
		Logger:                    zapLog,
		IgnoreRecordNotFoundError: false,
		traceStr:                  traceStr,
		traceWarnStr:              traceWarnStr,
		traceErrStr:               traceErrStr,
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  zapcore.WarnLevel,
	}
}
