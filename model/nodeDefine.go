package model

type NodeDefine struct {
	Id int `json:"id" gorm:"column:id"`
	// 所属 链 ID
	BelongLinkDefineId int `json:"belongLinkDefineId" gorm:"column:belong_link_define_id"`
	// 方法名
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 是否是链, 如果不是链, 那么就是节点, 可以是链, 也可以是节点, 如果是链的话,对应记录下 链的ID, 如果是节点,那么记录节点的生成代码实现方式
	IsLink bool `json:"isLink" gorm:"column:is_link"`
	// 如果是链, 那么需要记录 link define 的ID
	LinkDefineId int `json:"linkDefineId" gorm:"column:link_define_id"`
	// 前一个NodeID, 第一个写-1
	PreNodeDefineId int `json:"preNodeDefineId" gorm:"column:pre_node_define_id"`
}

func (receiver NodeDefine) TableName() string {
	return "node_define"
}

type NodeParamDefine struct {
	Id int `json:"id" gorm:"column:id"`
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

func (receiver NodeParamDefine) TableName() string {
	return "node_param_define"
}
