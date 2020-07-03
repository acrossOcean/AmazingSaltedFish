package model

type DBField struct {
	// 字段ID
	Id int `json:"fieldId" gorm:"column:id;primary_key"`
	// 对应 父级 ID (可能还是字段, 也可能是结构体)
	ParentId int `json:"parentId" gorm:"column:parent_id"`
	// 字段类型
	FType FieldType `json:"fieldType" gorm:"column:field_type" swaggertype:"integer"`
	// 字段名
	Name string `json:"name" gorm:"column:name"`
	// 字段注释
	Comment string `json:"comment" gorm:"column:comment"`
	// 是否为 列表
	IsList bool `json:"isList" gorm:"column:is_list"`
	// 排序位置
	Sort int `json:"sort" gorm:"column:sort"`

	// 对应结构体ID, 当 FType 为 FieldTypeStruct 时 有用
	StructId int `json:"structId" gorm:"column:struct_id"`

	// 当fieldType 为 map 时
	MapKeyFieldId   int `json:"mapKeyFieldId" gorm:"column:map_key_field_id"`
	MapValueFieldId int `json:"mapValueFieldId" gorm:"column:map_value_field_id"`

	// 如果是固定值, 那么记录
	IsConst   bool   `json:"isConst" gorm:"column:is_const"`
	ConstData []byte `json:"constData" gorm:"type:binary;column:const_data"`
}

func (receiver DBField) TableName() string {
	return "field_info"
}

type FieldList []DBField

func (receiver FieldList) Len() int { return len(receiver) }

func (receiver FieldList) Less(i, j int) bool {
	if receiver[i].Sort > receiver[j].Sort {
		return true
	} else if receiver[i].Sort < receiver[j].Sort {
		return false
	}

	return receiver[i].Id < receiver[j].Id
}

func (receiver FieldList) Swap(i, j int) { receiver[i], receiver[j] = receiver[j], receiver[i] }

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
	FieldTypeTime
	FieldTypeStruct
	FieldTypeMap
	FieldTypeInterface
)

func GetAllFieldType() []FieldType {
	result := []FieldType{
		FieldTypeString,
		FieldTypeBool,
		FieldTypeInt,
		FieldTypeFloat,
		FieldTypeTime,
		FieldTypeStruct,
		FieldTypeMap,
		FieldTypeInterface,
	}

	return result
}

func GetAllFieldTypeInt() []int {
	list := GetAllFieldType()
	result := make([]int, len(list))
	for i, ft := range list {
		result[i] = ft.ToInt()
	}

	return result
}
