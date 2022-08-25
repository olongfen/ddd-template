package errorx

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ BizError = (*bizError)(nil)

type BizError interface {
	// i 为了避免被其他包实现
	i()

	// WithError 设置错误信息
	WithError(err error) BizError

	// Code 获取业务码
	Code() BizCode

	Error() string

	// StackError 获取带堆栈的错误信息
	StackError() error
}

type bizError struct {
	code    BizCode // 业务码
	message string  // 错误描述
	stack   error   // 含有堆栈信息的错误
}

func NewError(bizCode BizCode) BizError {
	return &bizError{
		code:    bizCode,
		message: Text(bizCode),
	}
}

func (e *bizError) i() {}

func (e *bizError) Error() string {
	return fmt.Sprintf(`%s`, e.message)
}

func (e *bizError) WithError(err error) BizError {
	e.stack = errors.WithStack(err)
	return e
}

func (e *bizError) Code() BizCode {
	return e.code
}

func (e *bizError) Message() string {
	return e.message
}

func (e *bizError) StackError() error {
	return e.stack
}

func HandlerRecordNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewError(RecordNotFound).WithError(err)
	}
	return err
}
