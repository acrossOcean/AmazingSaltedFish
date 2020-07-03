package model

// DBLinkDefine link define 对应数据库结构
type DBLinkDefine struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 所属工程ID
	ProjectDefineID int `json:"projectDefineId" gorm:"column:project_define_id"`
	// 链名称
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 是否共享, 不共享的话 只能当前工程可用
	IsShared bool `json:"isShared" gorm:"column:is_shared"`
}

// TableName link define 对应数据库表名
func (receiver DBLinkDefine) TableName() string {
	return "link_define"
}

// DBLinkParamDefine link param define 对应数据库结构
type DBLinkParamDefine struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 所属 链 id
	LinkDefineID int `json:"linkDefineId" gorm:"column:link_define_id"`
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

// TableName link param define 对应数据库表名
func (receiver DBLinkParamDefine) TableName() string {
	return "link_param_define"
}
