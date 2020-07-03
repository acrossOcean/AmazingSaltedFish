package model

// 链 定义
type DBLinkInstance struct {
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 所属工程ID
	ProjectInstanceId int `json:"projectInstanceId" gorm:"column:project_instance_id"`
	// 所用 链 定义 id
	DefineId int `json:"defineId" gorm:"column:define_id"`
	// 链名称
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`

	// 是否共享, 不共享的话 只能当前工程可用
	IsShared bool `json:"isShared" gorm:"column:is_shared"`
}

func (receiver DBLinkInstance) TableName() string {
	return "link_instance"
}

type DBLinkParamInstance struct {
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 所用 链参数 定义 id
	DefineId int `json:"defineId" gorm:"column:define_id"`
	// 所属 链 instance id
	LinkInstanceId int `json:"linkInstanceId" gorm:"column:link_instance_id"`
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
	// 如果来自变量, 那么记录是来自 其它链还是其它Node
	InputVarIsLink bool `json:"inputVarIsLink" gorm:"column:input_var_is_link"`
	// 如果来自变量, 那么记录来源
	InputVarDefineId int `json:"inputVarDefineId" gorm:"column:input_var_define_id"`
	// 如果来自变量, 那么记录来源
	InputVarInstanceId int `json:"inputVarInstanceId" gorm:"column:input_var_instance_id"`

	// 返回值类型 (如果返回变量, 那么不需要记录, 到时候直接用就可以)
	OutputType ParamType `json:"outputType" gorm:"column:output_type"`
	// 如果是确定值, 那么记录这个值
	OutputConst []byte `json:"outputConst" gorm:"type:binary;column:output_const"`
}

func (receiver DBLinkParamInstance) TableName() string {
	return "link_param_instance"
}
