package dependency

import (
	"context"
	"ddd-template/domain/entities"
)

type DemoInterface interface {
	SayHello(ctx context.Context, msg string) *entities.Demo
}
