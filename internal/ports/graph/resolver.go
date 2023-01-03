package graph

import (
	app "ddd-template/internal/application"
	"github.com/99designs/gqlgen/graphql/handler"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app    *app.Application
	logger *zap.Logger
}

func NewResolver(app *app.Application, logger *zap.Logger) *handler.Server {
	rsl := &Resolver{app: app, logger: logger}
	c := Config{Resolvers: rsl}
	srv := handler.NewDefaultServer(NewExecutableSchema(c))
	return srv
}
