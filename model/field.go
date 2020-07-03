package model

// DBField field 对应 数据库结构
type DBField struct {
	// 字段ID
	ID int `json:"fieldId" gorm:"column:id;primary_key"`
	// 对应 父级 ID (可能还是字段, 也可能是结构体)
	ParentID int `json:"parentId" gorm:"column:parent_id"`
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
	StructID int `json:"structId" gorm:"column:struct_id"`

	// 当fieldType 为 map 时
	MapKeyFieldID   int `json:"mapKeyFieldId" gorm:"column:map_key_field_id"`
	MapValueFieldID int `json:"mapValueFieldId" gorm:"column:map_value_field_id"`

	// 如果是固定值, 那么记录
	IsConst   bool   `json:"isConst" gorm:"column:is_const"`
	ConstData []byte `json:"constData" gorm:"type:binary;column:const_data"`
}

// TableName field 对应数据库表名
func (receiver DBField) TableName() string {
	return "field_info"
}

// FieldList field 列表, 做排序使用
type FieldList []DBField

func (receiver FieldList) Len() int { return len(receiver) }

func (receiver FieldList) Less(i, j int) bool {
	if receiver[i].Sort > receiver[j].Sort {
		return true
	} else if receiver[i].Sort < receiver[j].Sort {
		return false
	}

	return receiver[i].ID < receiver[j].ID
}

func (receiver FieldList) Swap(i, j int) { receiver[i], receiver[j] = receiver[j], receiver[i] }

// FieldType field 类型
type FieldType int

// ToInt 供检查器使用.返回type的int值
func (receiver FieldType) ToInt() int {
	return int(receiver)
}

const (
	_ FieldType = iota
	// FieldTypeString :字段类型: String
	FieldTypeString
	// FieldTypeBool :字段类型: Bool
	FieldTypeBool
	// FieldTypeInt :字段类型: Int
	FieldTypeInt
	// FieldTypeFloat :字段类型: Float
	FieldTypeFloat
	// FieldTypeTime :字段类型: Time
	FieldTypeTime
	// FieldTypeStruct :字段类型: Struct
	FieldTypeStruct
	// FieldTypeMap :字段类型: Map
	FieldTypeMap
	// FieldTypeInterface :字段类型: Interface
	FieldTypeInterface
)

// GetAllFieldType 获取所有 支持的 field 类型
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

// GetAllFieldTypeInt 获取所有 支持的 field 类型 的 int 形式
func GetAllFieldTypeInt() []int {
	list := GetAllFieldType()
	result := make([]int, len(list))
	for i, ft := range list {
		result[i] = ft.ToInt()
	}

	return result
}
