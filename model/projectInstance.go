package model

// DBProjectInstance project instance 对应数据库结构
type DBProjectInstance struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 对应 定义 id
	DefineID int `json:"defineId" gorm:"column:define_id"`
	// 工程名
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`
	// 实现工程所用语言及版本
	GenerateLanguage string `json:"generateLanguage" gorm:"column:generate_language"`

	// 第一个 link 信息
	FirstLinkDefineID   int `json:"firstLinkDefineId" gorm:"column:first_link_define_id"`
	FirstLinkInstanceID int `json:"firstLinkInstanceId" gorm:"column:first_link_instance_id"`
}

// TableName project instance 对应数据库表名
func (receiver *DBProjectInstance) TableName() string {
	return "project_instance"
}

// DBProjectParamInstance project param instance 对应数据库结构
type DBProjectParamInstance struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 所属节点id
	ProjectInstanceID int `json:"projectInstanceId" gorm:"column:project_instance_id"`
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
	InputVarDefineID int `json:"inputVarDefineId" gorm:"column:input_var_define_id"`
	// 如果来自变量, 那么记录来源
	InputVarInstanceID int `json:"inputVarInstanceId" gorm:"column:input_var_instance_id"`

	// 返回值类型 (如果返回变量, 那么不需要记录, 到时候直接用就可以)
	OutputType ParamType `json:"outputType" gorm:"column:output_type"`
	// 如果是确定值, 那么记录这个值
	OutputConst []byte `json:"outputConst" gorm:"type:binary;column:output_const"`
}

// TableName project param instance 对应数据库表名
func (receiver DBProjectParamInstance) TableName() string {
	return "project_param_instance"
}
