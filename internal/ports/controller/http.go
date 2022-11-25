package controller

import (
	_ "ddd-template/docs"
	"ddd-template/internal/config"
	"ddd-template/internal/ports/controller/handler"
	"ddd-template/internal/ports/controller/middleware"
	"ddd-template/pkg/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
	"log"
)

type HttpServer struct {
	handler    *handler.Handler
	app        *fiber.App
	middleware middleware.Middleware
}

func (h *HttpServer) Close() (err error) {
	log.Println("http handler close")
	return h.app.Shutdown()
}

func NewHttpServer(handler *handler.Handler, m middleware.Middleware) *HttpServer {
	return &HttpServer{handler: handler, middleware: m}
}

func (h *HttpServer) RunHTTPServer(fc func(app2 *fiber.App) *fiber.App, cfg config.HTTP, logger *zap.Logger) {
	h.app = fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
		JSONEncoder:  jsoniter.Marshal,
		JSONDecoder:  jsoniter.Unmarshal,
	})
	h.app.Use(h.middleware.Languages(), h.middleware.Log())
	h.app.Mount("/api/v1", fc(h.app))
	logger.Info("HTTP Start", zap.String("addr", fmt.Sprintf(`%s:%d`, cfg.Host, cfg.Port)))
	logger.Fatal("HTTP START ERROR", zap.Error(h.app.Listen(fmt.Sprintf(`%s:%d`, cfg.Host, cfg.Port))))

}

func (h *HttpServer) HandlerFromMux(a *fiber.App) *fiber.App {
	a.Get("/docs/*", fiberSwagger.WrapHandler)
	return a
}
