package model

type Field struct {
	// 字段ID
	FId int `json:"fieldId" gorm:"column:id"`
	// 对应结构体ID
	ParentId int `json:"parentId" gorm:"column:parent_id"`
	// 字段类型
	FType FieldType `json:"fieldType" gorm:"column:field_type" swaggertype:"integer"`
	// 字段名
	FName string `json:"fieldName" gorm:"column:field_name"`
	// 字段注释
	FComment string `json:"fieldComment" gorm:"column:field_comment"`
	// 对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用
	FStructId int `json:"fieldStructId" gorm:"column:field_struct_id"`
	// 是否为 列表
	FIsList bool `json:"fieldIsList" gorm:"column:field_is_list"`
	// 排序位置
	FSort int `json:"fieldSort" gorm:"column:field_sort"`
}

func (receiver Field) TableName() string {
	return "field_info"
}

type FieldList []Field

func (receiver FieldList) Len() int { return len(receiver) }

func (receiver FieldList) Less(i, j int) bool {
	if receiver[i].FSort > receiver[j].FSort {
		return true
	} else if receiver[i].FSort < receiver[j].FSort {
		return false
	}

	return receiver[i].FId < receiver[j].FId
}

func (receiver FieldList) Swap(i, j int) { receiver[i], receiver[j] = receiver[j], receiver[i] }

type CreateFieldInfo struct {
	// 字段类型
	FType FieldType `json:"fieldType" swaggertype:"integer"`
	// 字段名
	FName string `json:"fieldName"`
	// 字段注释
	FComment string `json:"fieldComment"`
	// 对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用
	FStructId int `json:"fieldStructId"`
	// 是否为 列表
	FIsList bool `json:"fieldIsList"`
	// 排序位置
	FSort int `json:"fieldSort" gorm:"column:field_sort"`
}

func (receiver CreateFieldInfo) ToNormal() Field {
	var result = Field{
		FId:       0,
		ParentId:  0,
		FType:     receiver.FType,
		FName:     receiver.FName,
		FComment:  receiver.FComment,
		FStructId: receiver.FStructId,
		FIsList:   receiver.FIsList,
		FSort:     receiver.FSort,
	}

	return result
}

type FieldType int

func (receiver FieldType) ToInt() int {
	return int(receiver)
}

const (
	_ FieldType = iota
	FieldTypeString
	FieldTypeBool
	FieldTypeInt
	FieldTypeFloat
	FieldTypeStruct
)
