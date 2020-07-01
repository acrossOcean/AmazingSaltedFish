package model

type NodeInstance struct {
	Id int `json:"id" gorm:"column:id"`
	// 所属 链 ID
	BelongLinkInstanceId int `json:"belongLinkInstanceId" gorm:"column:belong_link_instance_id"`
	// 所用 node define id
	DefineId int `json:"defineId" gorm:"column:define_id"`
	// 方法名
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 是否是链, 如果不是链, 那么就是节点, 可以是链, 也可以是节点, 如果是链的话,对应记录下 链的ID, 如果是节点,那么记录节点的生成代码实现方式
	IsLink bool `json:"isLink" gorm:"column:is_link"`
	// 如果是链, 那么需要记录 link instance 的ID
	LinkInstanceId int `json:"linkInstanceId" gorm:"column:link_instance_id"`
	// 实现方式ID, 对应 "core_generator" id
	GeneratorId int `json:"generatorId" gorm:"column:generator_id"`
	// 实现方式名称
	GeneratorName string `json:"generatorName" gorm:"column:generator_name"`
	// 前一个NodeID, 第一个写-1
	PreNodeInstanceId int `json:"preNodeInstanceId" gorm:"column:pre_node_instance_id"`
}

func (receiver NodeInstance) TableName() string {
	return "node_instance"
}

type NodeParamInstance struct {
	Id int `json:"id" gorm:"column:id"`
	// 对应定义ID
	DefineId int `json:"defineId" gorm:"column:define_id"`
	// 所属节点id
	NodeInstanceId int `json:"nodeInstanceId" gorm:"column:node_instance_id"`
	// 位置, 1.入参 2.出参
	Location int `json:"isParam" gorm:"column:location"`
	// 字段类型
	PType FieldType `json:"type" gorm:"column:param_type" swaggertype:"integer"`
	// 字段名
	Name string `json:"name" gorm:"column:name"`
	// 字段注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 入参来源
	InputType ParamType `json:"inputType" gorm:"column:input_type"`
	// 如果是来自确定值, 那么记录这个值
	InputConst []byte `json:"inputConst" gorm:"type:binary;column:input_const"`
	// 如果来自变量, 那么记录是来自 其它link还是其它Node
	InputVarIsLink bool `json:"inputVarIsLink" gorm:"column:input_var_is_link"`
	// 如果来自变量, 那么记录来源
	InputVarDefineId int `json:"inputVarDefineId" gorm:"column:input_var_define_id"`
	// 如果来自变量, 那么记录来源
	InputVarInstanceId int `json:"inputVarInstanceId" gorm:"column:input_var_instance_id"`

	// 返回值类型 (如果返回变量, 那么不需要记录, 到时候直接用就可以)
	OutputType ParamType `json:"outputType" gorm:"column:output_type"`
	// 如果是确定值, 那么记录这个值
	OutputConst []byte `json:"outputConst" gorm:"type:binary;type:binary;column:output_const"`
}

func (receiver NodeParamInstance) TableName() string {
	return "node_param_instance"
}
