package model

type DBNodeDefine struct {
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 所属 链 ID
	BelongLinkDefineId int `json:"belongLinkDefineId" gorm:"column:belong_link_define_id"`
	// 方法名
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 前一个NodeID, 第一个写-1
	PreNodeDefineId int `json:"preNodeDefineId" gorm:"column:pre_node_define_id"`
}

func (receiver DBNodeDefine) TableName() string {
	return "node_define"
}

type DBNodeParamDefine struct {
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 所属节点id
	NodeDefineId int `json:"nodeDefineId" gorm:"column:node_define_id"`
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
	StructId int `json:"structId" gorm:"column:struct_id"`

	// 当ParamType 为 map 时
	MapKeyParamId   int `json:"mapKeyParamId" gorm:"column:map_key_param_id"`
	MapValueParamId int `json:"mapValueParamId" gorm:"column:map_value_param_id"`
}

func (receiver DBNodeParamDefine) TableName() string {
	return "node_param_define"
}
