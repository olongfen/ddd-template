// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package schema

type Demo struct {
	Msg string `json:"msg"`
}

type NewDemo struct {
	Msg string `json:"msg"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}
