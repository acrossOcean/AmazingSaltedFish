package model

// DBStructInfo 结构体 对应数据库结构
type DBStructInfo struct {
	// 结构体ID
	ID int `json:"id" gorm:"column:id;primary_key"`
	// 字段
	Fields []DBField `json:"fieldList" gorm:"column:-"`
	// 备注信息
	Comment string `json:"comment" gorm:"column:comment"`
}

// TableName 结构体 对应数据库 表名称
func (receiver DBStructInfo) TableName() string {
	return "struct_info"
}
