package errorx

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

	// WithAlert 设置告警通知
	WithAlert() BizError
	WithSupplementMsg(msg string) BizError

	// Code 获取业务码
	Code() BizCode

	// Message 获取错误描述
	Message() string

	Error() string

	// StackError 获取带堆栈的错误信息
	StackError() error

	// IsAlert 是否开启告警通知
	IsAlert() bool
}

type bizError struct {
	code    BizCode // 业务码
	message string  // 错误描述
	sup     string  // 补充的错误描述
	sta     error   // 含有堆栈信息的错误
	isAlert bool    // 是否告警通知
}

func NewError(bizCode BizCode) BizError {
	return &bizError{
		code:    bizCode,
		message: Text(bizCode),
		isAlert: false,
	}
}

func (e *bizError) i() {}

func (e *bizError) Error() string {
	if e.sta != nil {
		return fmt.Sprintf(`sta: %s\n biz: %s\n sup: %s\n`, e.sta.Error(), e.message, e.sup)
	}
	return fmt.Sprintf(`biz: %s\n sup: %s\n`, e.message, e.sup)
}

func (e *bizError) WithSupplementMsg(msg string) BizError {
	e.sup = msg
	return e
}

func (e *bizError) WithError(err error) BizError {
	e.sta = errors.WithStack(err)
	return e
}

func (e *bizError) WithAlert() BizError {
	e.isAlert = true
	return e
}

func (e *bizError) Code() BizCode {
	return e.code
}

func (e *bizError) Message() string {
	return e.message
}

func (e *bizError) StackError() error {
	return e.sta
}

func (e *bizError) IsAlert() bool {
	return e.isAlert
}
