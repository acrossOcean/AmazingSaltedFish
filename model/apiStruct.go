package model

// GetStructResp 获取 结构体 信息返回结构
type GetStructResp struct {
	BaseResp
	// 结构体信息
	Info DBStructInfo `json:"info"`
}

// CreateStructReq 新建 结构体 信息请求结构
type CreateStructReq struct {
	// 字段信息
	Fields []CreateFieldInfo `json:"fieldList"`
	// 字段注释
	Comment string `json:"comment"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver *CreateStructReq) ToDBStruct() DBStructInfo {
	var result = DBStructInfo{
		ID:      0,
		Fields:  make([]DBField, len(receiver.Fields)),
		Comment: receiver.Comment,
	}

	for i, f := range receiver.Fields {
		result.Fields[i] = f.ToDBStruct()
	}

	return result
}

// CreateStructResp 新建 结构体 信息返回结构
type CreateStructResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateStructReq 更新 结构体 信息请求结构
type UpdateStructReq struct {
	DBStructInfo
}

// UpdateStructResp 更新 结构体 信息返回结构
type UpdateStructResp struct {
	BaseResp
	// 更新的结构体ID
	ID int `json:"id"`
}

// DeleteStructResp 删除 结构体 信息返回结构
type DeleteStructResp struct {
	BaseResp
	// 删除的结构体ID
	ID int `json:"id"`
}

// GetStructListResp 获取 结构体 列表返回结构
type GetStructListResp struct {
	BaseResp
	// 结构体列表
	List []DBStructInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}
