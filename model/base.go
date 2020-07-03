package model

// ReturnCode 返回码
type ReturnCode int

const (
	// CodeSuccess : 请求成功
	CodeSuccess ReturnCode = iota
	// CodeUnknownError : 错误:未知错误
	CodeUnknownError
	// CodeParamError : 错误:传入参数错误
	CodeParamError
	// CodeResourceNotExist : 错误:请求资源没找到
	CodeResourceNotExist
)

var (
	// RespSuccess 成功返回的基础结构体
	RespSuccess = BaseResp{
		Code: CodeSuccess,
		Msg:  "success",
	}

	// RespParamError 参数错误返回的基础结构体
	RespParamError = BaseResp{
		Code: CodeParamError,
		Msg:  "输入参数错误",
	}
)

// PageReq 分页请求统一基础结构
type PageReq struct {
	// 页数
	PageNum int `json:"pageNum"`
	// 每页多少条信息
	PageSize int `json:"pageSize"`
}

// Format 格式化 分页请求, 返回合法的分页请求
func (receiver PageReq) Format() PageReq {
	var result = receiver
	if result.PageNum < 1 {
		result.PageNum = 1
	}

	if result.PageSize < 0 || result.PageSize > 500 {
		result.PageSize = 20
	}

	return result
}

// BaseResp 基础返回结构, 所有返回结构都包含此基础结构
type BaseResp struct {
	// 结果码
	Code ReturnCode
	// 提示信息
	Msg string
}

// SetSuccess 操作成功, 设置基础结构为 成功
func (receiver *BaseResp) SetSuccess() {
	receiver.Code = RespSuccess.Code
	receiver.Msg = RespSuccess.Msg
}

// SetError 操作失败, 设置基础结构包含错误信息
func (receiver *BaseResp) SetError(err error) {
	temp := NewBaseRespWithError(err)

	receiver.Code = temp.Code
	receiver.Msg = temp.Msg
}

// NewBaseRespWithError 根据传入 错误 生成新的基础结构
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
