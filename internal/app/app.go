package app

import (
	_ "ddd-template/doc"
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/response"
	"ddd-template/internal/service/delivery/xfiber"
	"ddd-template/internal/service/delivery/xfiber/middleware"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
	"sync"
)

type application struct {
	cfg      *conf.Configs
	logger   *zap.Logger
	handlers []IHandler
}

func NewApp(cfg *conf.Configs, logger *zap.Logger, demo *xfiber.DemoHandler) (err error) {
	app := &application{
		cfg:    cfg,
		logger: logger,
	}
	app.handlers = append(app.handlers, demo)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.startHTTP()
	}()
	wg.Wait()
	return
}

func (a *application) startHTTP() {
	cfg := fiber.Config{}
	cfg.JSONEncoder = jsoniter.Marshal
	cfg.JSONDecoder = jsoniter.Unmarshal
	app := fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
	})
	app.Use(middleware.Languages, middleware.New(middleware.Config{Logger: a.logger}))
	v1 := app.Group("/api/v1")
	v1.Get("/docs/*", fiberSwagger.WrapHandler)
	for _, handler := range a.handlers {
		handler.Handler(v1)
	}
	a.logger.Fatal("HTTP Server", zap.Error(app.Listen(a.cfg.HTTP.Addr)))
}
