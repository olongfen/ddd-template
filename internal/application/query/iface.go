package query

import "context"

// IDemoService demo query service
type IDemoService interface {
	Hello(ctx context.Context, msg string) string
}
