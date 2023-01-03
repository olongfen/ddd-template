package scontext

import (
	"context"
)

type languageCtxTag struct {
}

func SetLanguage(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, languageCtxTag{}, lang)
}

func GetLanguage(ctx context.Context) string {
	if val, ok := ctx.Value(languageCtxTag{}).(string); ok {
		return val
	}
	return "zh"
}

type userUuidCtxTag struct{}

func SetUserUuid(ctx context.Context, userUuid string) context.Context {
	return context.WithValue(ctx, userUuidCtxTag{}, userUuid)
}

func GetUserUuid(ctx context.Context) string {
	if val, ok := ctx.Value(userUuidCtxTag{}).(string); ok {
		return val
	}
	return ""
}
