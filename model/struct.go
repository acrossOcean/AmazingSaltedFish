package model

type StructInfo struct {
	// 结构体ID
	Id int `json:"id" gorm:"column:id"`
	// 字段
	Fields []Field `json:"fieldList" gorm:"column:-"`
	// 备注信息
	Comment string `json:"comment" gorm:"column:comment"`
}

func (receiver StructInfo) TableName() string {
	return "struct_info"
}

type GetStructResp struct {
	BaseResp
	// 结构体信息
	Info StructInfo `json:"info"`
}

type CreateStructReq struct {
	// 字段信息
	Fields []CreateFieldInfo `json:"fieldList"`
	// 字段注释
	Comment string `json:"comment"`
}

func (receiver *CreateStructReq) ToNormal() StructInfo {
	var result = StructInfo{
		Id:      0,
		Fields:  make([]Field, len(receiver.Fields)),
		Comment: receiver.Comment,
	}

	for i, f := range receiver.Fields {
		result.Fields[i] = f.ToNormal()
	}

	return result
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
	// 总数据条数
	SUm int `json:"sum"`
}
