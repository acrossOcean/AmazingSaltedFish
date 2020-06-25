package model

type PageReq struct {
	// 页数
	PageNum int `json:"pageNum"`
	// 每页多少条信息
	PageSize int `json:"pageSize"`
}

type BaseResp struct {
	// 结果码
	Code string
	// 提示信息
	Msg string
}
