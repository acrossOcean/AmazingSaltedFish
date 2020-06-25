package model

type StructInfo struct {
	// 结构体ID
	Id int `json:"id"`
	// 字段
	Fields []Field `json:"fieldList"`
	// 备注信息
	Comment string `json:"comment"`
}

type GetStructResp struct {
	BaseResp
	// 结构体信息
	Info StructInfo `json:"info"`
}

type CreateStructReq struct {
	// 字段信息
	Fields []Field `json:"fieldList"`
	// 字段注释
	Comment string `json:"comment"`
}

type CreateStructResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateStructReq struct {
	StructInfo
}

type UpdateStructResp struct {
	BaseResp
	// 更新的结构体ID
	Id int `json:"id"`
}

type DeleteStructResp struct {
	BaseResp
	// 删除的结构体ID
	Id int `json:"id"`
}

type GetStructListResp struct {
	BaseResp
	// 结构体列表
	List []StructInfo `json:"list"`
}
