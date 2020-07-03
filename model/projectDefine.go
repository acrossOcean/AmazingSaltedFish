package model

// DBProjectDefine project define 对应数据库结构
type DBProjectDefine struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 工程名
	Name string `json:"name" gorm:"column:name"`
	// 注释
	Comment string `json:"comment" gorm:"column:comment"`
}

// TableName project define 对应数据库表名
func (receiver DBProjectDefine) TableName() string {
	return "project_define"
}

// DBProjectParamDefine project param define 对应数据库结构
type DBProjectParamDefine struct {
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 所属项目id
	ProjectDefineID int `json:"projectDefineId" gorm:"column:project_define_id"`
	// 位置, 1.入参 2.出参
	Location int `json:"isParam" gorm:"column:location"`
	// 字段类型
	PType ParamType `json:"type" gorm:"column:param_type" swaggertype:"integer"`
	// 字段名
	Name string `json:"name" gorm:"column:name"`
	// 字段注释
	Comment string `json:"comment" gorm:"column:comment"`
	// 是否为 列表
	IsList bool `json:"isList" gorm:"column:is_list"`
	// 排序位置
	Sort int `json:"sort" gorm:"column:sort"`
}

// TableName project param define 对应数据库表名
func (receiver DBProjectParamDefine) TableName() string {
	return "project_param_define"
}
