package model

// CreateFieldInfo 新建 field 信息请求结构
type CreateFieldInfo struct {
	// 字段类型
	Type FieldType `json:"type" swaggertype:"integer"`
	// 字段名
	Name string `json:"name"`
	// 字段注释
	Comment string `json:"comment"`
	// 对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用
	StructID int `json:"structId"`
	// 是否为 列表
	IsList bool `json:"isList"`
	// 排序位置
	Sort int `json:"sort"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateFieldInfo) ToDBStruct() DBField {
	var result = DBField{
		ID:       0,
		ParentID: 0,
		FType:    receiver.Type,
		Name:     receiver.Name,
		Comment:  receiver.Comment,
		StructID: receiver.StructID,
		IsList:   receiver.IsList,
		Sort:     receiver.Sort,
	}

	return result
}
