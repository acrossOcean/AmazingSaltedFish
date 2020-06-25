package model

type FieldType int

type Field struct {
	// 字段类型
	FType FieldType `json:"fieldType" swaggertype:"integer"`
	// 字段名
	FName string `json:"fieldName"`
	// 字段注释
	FComment string `json:"fieldComment"`
	// 对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用
	FStructId int `json:"fieldStructId"`
	// 默认值
	FDefault string `json:"fieldDefaultValue"`
	// 是否为 列表
	FIsList bool `json:"fieldIsList"`
}

const (
	_ FieldType = iota
	FieldTypeString
	FieldTypeBool
	FieldTypeInt
	FieldTypeFloat
	FieldTypeStruct
)
