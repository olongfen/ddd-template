package xi18n

import (
	"fmt"
	"github.com/pkg/errors"
)

var _ BizError = (*bizError)(nil)

type BizError interface {
	// i 为了避免被其他包实现
	i()

	// WithError 设置错误信息
	WithError(err error) BizError

	Error() string

	// StackError 获取带堆栈的错误信息
	StackError() error
}

type bizError struct {
	message string // 错误描述
	stack   error  // 含有堆栈信息的错误
}

func NewError(msg string) BizError {
	biz := &bizError{
		message: msg,
	}
	return biz
}

func (e *bizError) i() {}

func (e *bizError) Error() string {
	return fmt.Sprintf(`%s`, e.message)
}

func (e *bizError) WithError(err error) BizError {
	e.stack = errors.WithStack(err)
	return e
}

func (e *bizError) Message() string {
	return e.message
}

func (e *bizError) StackError() error {
	return e.stack
}
