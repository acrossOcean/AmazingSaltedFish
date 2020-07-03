package model

type GetStructResp struct {
	BaseResp
	// 结构体信息
	Info DBStructInfo `json:"info"`
}

type CreateStructReq struct {
	// 字段信息
	Fields []CreateFieldInfo `json:"fieldList"`
	// 字段注释
	Comment string `json:"comment"`
}

func (receiver *CreateStructReq) ToDBStruct() DBStructInfo {
	var result = DBStructInfo{
		Id:      0,
		Fields:  make([]DBField, len(receiver.Fields)),
		Comment: receiver.Comment,
	}

	for i, f := range receiver.Fields {
		result.Fields[i] = f.ToDBStruct()
	}

	return result
}

type CreateStructResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateStructReq struct {
	DBStructInfo
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
	List []DBStructInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}
