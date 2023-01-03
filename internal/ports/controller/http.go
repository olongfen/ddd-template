package controller

import (
	_ "ddd-template/docs"
	"ddd-template/internal/config"
	"ddd-template/internal/ports/controller/handler"
	"ddd-template/internal/ports/controller/middleware"
	"ddd-template/pkg/response"
	"fmt"
	handler2 "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type HttpServer struct {
	handler        *handler.Handler
	app            *fiber.App
	middleware     middleware.Middleware
	graphqlHandler *handler2.Server
}

func (h *HttpServer) Close() (err error) {
	log.Println("http handler close")
	return h.app.Shutdown()
}

func NewHttpServer(handler *handler.Handler, graphqlHandler *handler2.Server, m middleware.Middleware) *HttpServer {
	return &HttpServer{handler: handler, middleware: m, graphqlHandler: graphqlHandler}
}

func (h *HttpServer) RunHTTPServer(fc func(app2 *fiber.App) *fiber.App, cfg config.HTTP, logger *zap.Logger) {
	h.app = fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
		JSONEncoder:  jsoniter.Marshal,
		JSONDecoder:  jsoniter.Unmarshal,
	})
	h.app.Use(h.middleware.Languages(), h.middleware.Log())
	h.app.Mount("/api/v1", fc(h.app))
	logger.Info("HTTP Start", zap.String("addr", fmt.Sprintf(`%s:%s`, cfg.Host, cfg.Port)))
	logger.Fatal("HTTP START ERROR", zap.Error(h.app.Listen(fmt.Sprintf(`%s:%s`, cfg.Host, cfg.Port))))

}

func (h *HttpServer) HandlerFromMux(a *fiber.App) *fiber.App {
	a.Get("/docs/*", fiberSwagger.WrapHandler)
	groupGraphql := a.Group("/")
	groupGraphql.All("/", HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	groupGraphql.All("/query", HTTPHandler(h.graphqlHandler))
	return a
}

func HTTPHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fastHTTPHandler := fasthttpadaptor.NewFastHTTPHandler(h)
		fastHTTPHandler(c.Context())
		return nil
	}
}
