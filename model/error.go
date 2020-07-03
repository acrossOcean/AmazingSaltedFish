package model

import "fmt"

// CError 自定义错误类型, 包含上 msg 信息
type CError struct {
	Code ReturnCode
	Msg  string
}

func (receiver CError) Error() string {
	return receiver.String()
}

func (receiver CError) String() string {
	return fmt.Sprintf("error(%d): %s", receiver.Code, receiver.Msg)
}

// ReturnWithCode 便捷返回 CError
func ReturnWithCode(code ReturnCode, err error) error {
	return CError{
		Code: CodeResourceNotExist,
		Msg:  err.Error(),
	}
}
