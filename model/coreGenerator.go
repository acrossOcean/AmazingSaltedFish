package model

// DBCoreGenerator 实现核心, 所有核心的逻辑全放到这里
// 对应数据库结构
type DBCoreGenerator struct {
	ID int `json:"id" gorm:"column:id;primary_key"`

	// 实现所用语言
	Language string `json:"language" gorm:"column:language"`
	// 名字
	Name string `json:"name" gorm:"column:name"`
	// Version. 实现对应版本
	Version string `json:"version" gorm:"column:version"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 实现内容
	Data string `json:"-" gorm:"type:text;column:data"`
}

// TableName 对应数据库表名
func (receiver *DBCoreGenerator) TableName() string {
	return "core_generator"
}

// DBCoreGeneratorParam 实现核心所需参数
// 对应数据库结构
type DBCoreGeneratorParam struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 所属生成器id
	GeneratorID int `json:"generatorId" gorm:"column:generator_id"`
	// 位置, 1.入参 2.出参
	Location int `json:"isParam" gorm:"column:location"`
	// 字段类型
	PType FieldType `json:"type" gorm:"column:param_type" swaggertype:"integer"`
	// 字段名
	Name string `json:"name" gorm:"column:name"`
	// 字段注释
	Comment string `json:"comment" gorm:"column:comment"`
	// 是否为 列表
	IsList bool `json:"isList" gorm:"column:is_list"`
	// 排序位置
	Sort int `json:"sort" gorm:"column:sort"`

	// 对应结构体ID, 当 FType 为 ParamTypeStruct 时 有用
	StructID int `json:"structId" gorm:"column:struct_id"`

	// 当ParamType 为 map 时
	MapKeyParamID   int `json:"mapKeyParamId" gorm:"column:map_key_param_id"`
	MapValueParamID int `json:"mapValueParamId" gorm:"column:map_value_param_id"`

	// 如果是固定值, 那么记录
	IsConst   bool   `json:"isConst" gorm:"column:is_const"`
	ConstData []byte `json:"constData" gorm:"type:binary;column:const_data"`
}

// TableName 对应数据库表名
func (receiver *DBCoreGeneratorParam) TableName() string {
	return "core_generator_param"
}
