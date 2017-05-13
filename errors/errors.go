package errors

import (
	"fmt"
	"strings"
)

// NewCodeError 新建一个error
func NewCodeError(code int, msgs ...interface{}) CodeError {
	return CodeError{Code: code, Msg: fmt.Sprintf(strings.Repeat("%s", len(msgs)), msgs...)}
}

// ConstError 不可变error
type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

// CodeError 带错误码的错误
type CodeError struct {
	Msg  string
	Code int
}

func (e CodeError) Error() string {
	return fmt.Sprint(e.Msg, e.Code)
}
