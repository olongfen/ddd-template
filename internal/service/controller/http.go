package controller

import (
	_ "ddd-template/docs"
	"ddd-template/internal/rely"
	"ddd-template/internal/service/controller/handler"
	"ddd-template/internal/service/graph"
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
	middleware    Middleware
	graphResolver *graph.Resolver
	logger        *zap.Logger
	cfg           *rely.Configs
}

// NewHTTPServer new http server
func NewHTTPServer(handler *handler.Handler, graphResolver *graph.Resolver, m Middleware,
	cfg *rely.Configs,
	logger *zap.Logger) (*HTTPServer, func()) {
	h := &HTTPServer{handler: handler, middleware: m, graphResolver: graphResolver, logger: logger, cfg: cfg}
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
func (h *HTTPServer) Run() {
	// http restful
	h.app.Use(h.middleware.Languages(), h.middleware.Log())
	var v1 = h.app.Group("/api/v1")
	h.handler.Process(v1)
	v1.Get("/docs/*", fiberSwagger.WrapHandler)
	// graphql
	if h.cfg.EnableGraph {
		h.graphResolver.Process(h.app.Group("/"))
	}
	h.logger.Info("HTTP Start", zap.String("addr", fmt.Sprintf(`%s:%s`, h.cfg.HTTP.Host, h.cfg.HTTP.Port)))
	if err := h.app.Listen(fmt.Sprintf(`%s:%s`, h.cfg.HTTP.Host, h.cfg.HTTP.Port)); err != nil {
		log.Fatalln(err)
	}

}
