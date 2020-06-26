package model

import "fmt"

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

func ReturnWithCode(code ReturnCode, err error) error {
	return CError{
		Code: CodeResourceNotExist,
		Msg:  err.Error(),
	}
}
