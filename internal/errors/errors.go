package errors

import "fmt"

// Error 异常
type Error struct {
	format string
	args   []interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf(e.format, e.args...)
}

func (e *Error) String() string {
	return e.Error()
}

// Errorf 自定义参数
func (e *Error) Errorf(args ...interface{}) string {
	e.args = args
	return e.Error()
}

// New create a new Error
func New(format string, args ...interface{}) *Error {
	newErr := &Error{
		format: format,
		args:   args,
	}
	return newErr
}
