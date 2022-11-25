package scontext

import (
	"context"
)

type Error struct {
	Failed string      `json:"failed"`
	Tag    string      `json:"tag"`
	Value  interface{} `json:"value"`
	Detail string      `json:"detail"`
}

type languageCtxTag struct {
}

type errorsCtxTag struct {
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

func SetErrorsContext(ctx context.Context, val map[string]*Error) context.Context {
	return context.WithValue(ctx, errorsCtxTag{}, val)
}

func GetErrorsContext(ctx context.Context) map[string]*Error {
	if val, ok := ctx.Value(errorsCtxTag{}).(map[string]*Error); ok {
		return val
	}
	return map[string]*Error{}
}
