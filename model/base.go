package model

type ReturnCode int

const (
	CodeSuccess ReturnCode = iota
	CodeUnknownError
	CodeParamError
	CodeResourceNotExist
)

var (
	RespSuccess = BaseResp{
		Code: CodeSuccess,
		Msg:  "",
	}

	RespParamError = BaseResp{
		Code: CodeParamError,
		Msg:  "输入参数错误",
	}
)

type PageReq struct {
	// 页数
	PageNum int `json:"pageNum"`
	// 每页多少条信息
	PageSize int `json:"pageSize"`
}

type BaseResp struct {
	// 结果码
	Code ReturnCode
	// 提示信息
	Msg string
}

func (receiver *BaseResp) SetSuccess() {
	receiver.Code = CodeSuccess
	receiver.Msg = "success"
}

func (receiver *BaseResp) SetError(err error) {
	temp := NewBaseRespWithError(err)

	receiver.Code = temp.Code
	receiver.Msg = temp.Msg
}

func NewBaseRespWithError(err error) BaseResp {
	switch err := err.(type) {
	case CError:
		return BaseResp(err)
	default:
		return BaseResp{
			Code: CodeUnknownError,
			Msg:  err.Error(),
		}
	}
}
