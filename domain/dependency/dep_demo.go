package dependency

import "context"

type DemoInterface interface {
	SayHello(ctx context.Context,msg string) string
}
