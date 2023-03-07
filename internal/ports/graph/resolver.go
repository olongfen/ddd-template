package graph

import (
	app "ddd-template/internal/application"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
	"net/http"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app    *app.Application
	logger *zap.Logger
}

// NewResolver new
func NewResolver(app *app.Application, logger *zap.Logger) *Resolver {
	rsl := &Resolver{app: app, logger: logger}
	return rsl
}

// Process handler
func (r *Resolver) Process(group fiber.Router) {
	c := Config{Resolvers: r}
	srv := handler.NewDefaultServer(NewExecutableSchema(c))
	group.All("/", HTTPHandler2FastHTTPHandler(playground.Handler("GraphQL playground", "/query")))
	group.All("/query", HTTPHandler2FastHTTPHandler(srv))
}

// HTTPHandler2FastHTTPHandler http handler 2 fast http handler
func HTTPHandler2FastHTTPHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fastHTTPHandler := fasthttpadaptor.NewFastHTTPHandler(h)
		fastHTTPHandler(c.Context())
		return nil
	}
}
