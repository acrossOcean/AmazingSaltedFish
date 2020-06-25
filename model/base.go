package model

type PageReq struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type BaseResp struct {
	Code string
	Msg  string
}
