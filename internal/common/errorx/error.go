package errorx

import (
	"fmt"
	"github.com/pkg/errors"
)

var _ BusinessError = (*businessError)(nil)

type BusinessError interface {
	// i 为了避免被其他包实现
	i()

	// WithError 设置错误信息
	WithError(err error) BusinessError

	// WithAlert 设置告警通知
	WithAlert() BusinessError
	WithSupplementMsg(msg string) BusinessError

	// Code 获取业务码
	Code() int

	// Message 获取错误描述
	Message() string

	Error() string

	// StackError 获取带堆栈的错误信息
	StackError() error

	// IsAlert 是否开启告警通知
	IsAlert() bool
}

type businessError struct {
	code          int    // 业务码
	message       string // 错误描述
	supplementMsg string // 补充的错误描述
	stackError    error  // 含有堆栈信息的错误
	isAlert       bool   // 是否告警通知
}

func Error(businessCode int) BusinessError {
	return &businessError{
		code:    businessCode,
		message: Text(businessCode),
		isAlert: false,
	}
}

func (e *businessError) i() {}

func (e *businessError) Error() string {
	if e.stackError != nil {
		return fmt.Sprintf(`stack: %s, business: %s, supplement: %s`, e.stackError.Error(), e.message, e.supplementMsg)
	}
	return fmt.Sprintf(`business: %s,supplement: %s`, e.message, e.supplementMsg)
}

func (e *businessError) WithSupplementMsg(msg string) BusinessError {
	e.supplementMsg = msg
	return e
}

func (e *businessError) WithError(err error) BusinessError {
	e.stackError = errors.WithStack(err)
	return e
}

func (e *businessError) WithAlert() BusinessError {
	e.isAlert = true
	return e
}

func (e *businessError) Code() int {
	return e.code
}

func (e *businessError) Message() string {
	return e.message
}

func (e *businessError) StackError() error {
	return e.stackError
}

func (e *businessError) IsAlert() bool {
	return e.isAlert
}
