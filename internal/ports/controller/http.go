package controller

import (
	_ "ddd-template/docs"
	"ddd-template/internal/config"
	"ddd-template/internal/ports/controller/middleware"
	"ddd-template/pkg/response"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
)

type HttpServer struct {
	Server Server
}

func NewHttpServer(server Server) (HttpServer, func()) {
	return HttpServer{server}, server.Cleanup()
}

func RunHTTPServer(fc func(app2 *fiber.App) *fiber.App, cfg config.HTTP, logger *zap.Logger) {
	a := fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
		JSONEncoder:  jsoniter.Marshal,
		JSONDecoder:  jsoniter.Unmarshal,
	})
	a.Use(middleware.Languages, middleware.New(middleware.Config{Logger: logger}))
	a.Mount("/api/v1", fc(a))
	logger.Info("HTTP Start", zap.String("addr", cfg.Host+":"+cfg.Port))
	logger.Fatal("HTTP START ERROR", zap.Error(a.Listen(cfg.Host+":"+cfg.Port)))

}

func HandlerFromMux(server Server, a *fiber.App) *fiber.App {
	a.Get("/docs/*", fiberSwagger.WrapHandler)
	// student
	stu := a.Group("/students")
	stu.Post("/", server.AddStudent)
	stu.Get("/:id", server.GetStudent)
	stu.Get("/", server.QueryStudents)
	stu.Put("/:id", server.UpStudent)
	// class
	class := a.Group("/classes")
	class.Post("/", server.AddClass)
	return a
}

type Server interface {
	Cleanup() func()
	AddStudent(ctx *fiber.Ctx) error
	GetStudent(ctx *fiber.Ctx) error
	QueryStudents(ctx *fiber.Ctx) error
	UpStudent(ctx *fiber.Ctx) error
	//
	AddClass(ctx *fiber.Ctx) error
}
