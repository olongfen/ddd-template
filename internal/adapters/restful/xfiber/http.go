package xfiber

import (
	_ "ddd-template/api"
	v1 "ddd-template/api/v1"
	"ddd-template/internal/app"
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/xlog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jsoniter "github.com/json-iterator/go"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type HTTPServer struct {
	engine *fiber.App
	cfg    conf.HTTP
	demo   v1.GreeterServer
}

func NewHTTPServer(demo v1.GreeterServer, cfg *conf.Configs) app.HttpServer {
	h := new(HTTPServer)
	c := fiber.Config{
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	}
	h.cfg = cfg.Server.Http
	h.demo = demo
	h.engine = fiber.New(c)
	return h
}

func (h *HTTPServer) Run() error {
	h.engine.Use(cors.New())
	h.engine.Get("/api/v1/docs/*", fiberSwagger.WrapHandler)
	xlog.Log.Sugar().Infof("http server run in: %s", h.cfg.Addr)
	group1 := h.engine.Group("/")
	v1.RegisterGreeterFiberHTTPServer(group1, h.demo)
	return h.engine.Listen(h.cfg.Addr)
}
