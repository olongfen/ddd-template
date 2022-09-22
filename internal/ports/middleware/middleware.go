package middleware

import (
	"bytes"
	"ddd-template/pkg/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Config defines the config for middleware
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Logger defines zap logger instance
	Logger *zap.Logger
}

// New creates a new middleware handler
func New(config Config) fiber.Handler {
	var (
		errPadding  = 15
		start, stop time.Time
		once        sync.Once
		errHandler  fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) error {
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		once.Do(func() {
			errHandler = c.App().Config().ErrorHandler
			stack := c.App().Stack()
			for m := range stack {
				for r := range stack[m] {
					if len(stack[m][r].Path) > errPadding {
						errPadding = len(stack[m][r].Path)
					}
				}
			}
		})

		start = time.Now()

		chainErr := c.Next()

		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		stop = time.Now()

		fields := []zap.Field{
			zap.Namespace("context"),
			zap.String("pid", strconv.Itoa(os.Getpid())),
			zap.String("time", stop.Sub(start).String()),
			zap.Object("response", Resp(c.Response())),
			zap.Object("request", Req(c)),
		}

		if u := c.Locals("userId"); u != nil {
			fields = append(fields, zap.Uint("userId", u.(uint)))
		}

		formatErr := ""
		if chainErr != nil && c.Response().StatusCode() != fiber.StatusOK {
			formatErr = chainErr.Error()
			fields = append(fields, zap.String("error", formatErr))
			config.Logger.With(fields...).Error(formatErr)

			return nil
		}

		config.Logger.With(fields...).Info("api.request")

		return nil
	}
}

func getAllowedHeaders() map[string]bool {
	return map[string]bool{
		"User-Agent": true,
		"X-Mobile":   true,
	}
}

type resp struct {
	code  int
	_type string
}

func Resp(r *fasthttp.Response) *resp {
	return &resp{
		code:  r.StatusCode(),
		_type: bytes.NewBuffer(r.Header.ContentType()).String(),
	}
}

func (r *resp) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", r._type)
	enc.AddInt("code", r.code)

	return nil
}

type req struct {
	body     string
	fullPath string
	user     string
	ip       string
	method   string
	route    string
	headers  *headerBag
}

func Req(c *fiber.Ctx) *req {
	reqq := c.Request()
	var body []byte
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, reqq.Body())
	if err != nil {
		body = reqq.Body()
	} else {
		body = buffer.Bytes()
	}

	headers := &headerBag{
		values: make(map[string]string),
	}
	// 获取请求头部信息
	allowedHeaders := getAllowedHeaders()
	reqq.Header.VisitAll(func(key, val []byte) {
		k := bytes.NewBuffer(key).String()
		if _, exist := allowedHeaders[k]; exist {
			headers.values[strings.ToLower(k)] = bytes.NewBuffer(val).String()
		}
	})

	// 获取用户信息
	var userEmail string
	if u := c.Locals("userEmail"); u != nil {
		userEmail = u.(string)
	}

	return &req{
		body:     bytes.NewBuffer(body).String(),
		fullPath: bytes.NewBuffer(reqq.RequestURI()).String(),
		headers:  headers,
		ip:       c.IP(),
		method:   c.Method(),
		route:    c.Route().Path,
		user:     userEmail,
	}
}

func (r *req) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("fullPath", r.fullPath)
	enc.AddString("ip", r.ip)
	enc.AddString("method", r.method)
	enc.AddString("route", r.route)

	if r.body != "" {
		enc.AddString("body", r.body)
	}

	if r.user != "" {
		enc.AddString("user", r.user)
	}

	err := enc.AddObject("headers", r.headers)
	if err != nil {
		return err
	}

	return nil

}

type headerBag struct {
	values map[string]string
}

func (h *headerBag) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, v := range h.values {
		enc.AddString(k, v)
	}

	return nil
}

func Languages(ctx *fiber.Ctx) error {
	val := ctx.GetReqHeaders()["x-language"]
	if len(val) == 0 {
		val = "en"
	}
	ctx.SetUserContext(utils.SetLanguage(ctx.UserContext(), val))
	return ctx.Next()
}
