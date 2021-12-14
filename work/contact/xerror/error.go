package xerror

import (
	"strings"

	"github.com/pkg/errors"
)

// go:generate stringer -type=Error -linecomment -output error_string.go
// type Error int // Error 错误
//
// // Error 输出错误信息
// func (i Error) Error() string {
// 	return i.String()
// }

// Wrap 包装错误信息
func Wrap(err error, args ...interface{}) error {
	if len(args) >= 1 {
		if msg, ok := args[0].(string); ok {
			return errors.Wrap(err, msg)
		}
	}
	return errors.Wrap(err, "")
}

// Cause 获取原始错误对象
func Cause(err error) error {
	return errors.Cause(err)
}

// Errorf 创建新错误
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

type codeError struct {
	code int64
	err  string
}

func (c codeError) Error() string {
	return c.err
}

// NewSDKErr 新建业务错误，附带错误码
func NewSDKErr(code int64, msgList ...string) error {
	return codeError{code: code, err: strings.Join(msgList, ",")}
}

// Code 提取错误码，codeError 返回 code 和 true，其他返回 0 和 false
func Code(err error) (int64, bool) {
	err = errors.Cause(err)
	if err == nil {
		return 0, false
	}
	if ce, ok := err.(codeError); ok {
		return ce.code, true
	}
	return 0, false
}
