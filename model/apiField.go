package model

type CreateFieldInfo struct {
	// 字段类型
	Type FieldType `json:"type" swaggertype:"integer"`
	// 字段名
	Name string `json:"name"`
	// 字段注释
	Comment string `json:"comment"`
	// 对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用
	StructId int `json:"structId"`
	// 是否为 列表
	IsList bool `json:"isList"`
	// 排序位置
	Sort int `json:"sort"`
}

func (receiver CreateFieldInfo) ToDBStruct() DBField {
	var result = DBField{
		Id:       0,
		ParentId: 0,
		FType:    receiver.Type,
		Name:     receiver.Name,
		Comment:  receiver.Comment,
		StructId: receiver.StructId,
		IsList:   receiver.IsList,
		Sort:     receiver.Sort,
	}

	return result
}
