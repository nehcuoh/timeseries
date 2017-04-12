package face

import (
	"fmt"
	"runtime/debug"
)

type Error struct {
	code int
	msg  string
}

type FatalError struct {
	err   Error
	stack []byte
}

func (e *Error) Error() string {
	return e.msg + fmt.Sprintf("[ %d ]:", e.code)
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func NewFatalError(code int, msg string) *FatalError {
	return &FatalError{err: Error{code: code, msg: msg}, stack: debug.Stack() }
}

func (e *Error) Code() int {
	return e.code
}
func (e*Error) Msg() (string) {
	return e.msg
}
func (e*FatalError) Stack() string {
	return string(e.stack[:])
}
func (e *FatalError) Code() int {
	return e.err.Code()
}
func (e*FatalError) Msg() (string) {
	return e.err.Msg()
}
func (e*FatalError) Error() string {
	return e.err.Error() + e.Stack()
}

func (e*FatalError) SetStack(stack []byte) {
	e.stack=stack
}
