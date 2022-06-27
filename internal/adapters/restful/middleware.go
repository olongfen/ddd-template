package restful

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"olongfen/ddd-template/internal/common/xlog"
	"time"
)

func corsHandler() gin.HandlerFunc {
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	return cors.New(corsCfg)
}

type LogFormatter func(params gin.LogFormatterParams) map[string]interface{}

func format(f LogFormatter) gin.HandlerFunc {
	logClient := xlog.Log
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		logClient.Sugar().Info(f(param))
	}
}

type formatter struct {
	Address    string        `json:"address"`
	Time       string        `json:"time"`
	Protoc     string        `json:"protoc"`
	Method     string        `json:"method"`
	Path       string        `json:"path"`
	Code       int           `json:"code"`
	Body       any           `json:"body"`
	BodySize   int           `json:"bodySize"`
	Latency    time.Duration `json:"latency"`
	ErrMessage string        `json:"errMessage"`
}

func logger() gin.HandlerFunc {
	return format(
		func(params gin.LogFormatterParams) map[string]interface{} {
			var res = formatter{
				Address:    params.ClientIP,
				Time:       params.TimeStamp.Format("2006-01-02 15:04:05"),
				Protoc:     params.Request.Proto,
				Method:     params.Request.Method,
				Path:       params.Path,
				Code:       params.StatusCode,
				Body:       params.Request.Body,
				BodySize:   params.BodySize,
				Latency:    params.Latency,
				ErrMessage: params.ErrorMessage,
			}
			var (
				m = make(map[string]interface{})
			)
			bd, _ := jsoniter.Marshal(res)
			jsoniter.Unmarshal(bd, &m)
			return m
		})
}
