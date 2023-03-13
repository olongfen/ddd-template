package controller

import (
	_ "ddd-template/docs"
	"ddd-template/internal/ports/controller/handler"
	"ddd-template/internal/ports/controller/middleware"
	"ddd-template/internal/ports/graph"
	"ddd-template/internal/rely"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/olongfen/toolkit/response"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
	"log"
)

// HTTPServer http server
type HTTPServer struct {
	handler       *handler.Handler
	app           *fiber.App
	middleware    middleware.Middleware
	graphResolver *graph.Resolver
	logger        *zap.Logger
}

// NewHTTPServer new http server
func NewHTTPServer(handler *handler.Handler, graphResolver *graph.Resolver, m middleware.Middleware,
	logger *zap.Logger) (*HTTPServer, func()) {
	h := &HTTPServer{handler: handler, middleware: m, graphResolver: graphResolver, logger: logger}
	// new app
	h.app = fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
		JSONEncoder:  jsoniter.Marshal,
		JSONDecoder:  jsoniter.Unmarshal,
	})
	return h, func() {
		log.Println("http server close")
		_ = h.app.Shutdown()

	}
}

// Run run http server
func (h *HTTPServer) Run(cfg rely.HTTP) {
	// http restful
	h.app.Use(h.middleware.Languages(), h.middleware.Log())
	var v1 = h.app.Group("/api/v1")
	h.handler.Process(v1)
	v1.Get("/docs/*", fiberSwagger.WrapHandler)
	// graphql
	h.graphResolver.Process(h.app.Group("/"))
	h.logger.Info("HTTP Start", zap.String("addr", fmt.Sprintf(`%s:%s`, cfg.Host, cfg.Port)))
	h.logger.Fatal("HTTP START ERROR", zap.Error(h.app.Listen(fmt.Sprintf(`%s:%s`, cfg.Host, cfg.Port))))

}
