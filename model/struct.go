package model

type DBStructInfo struct {
	// 结构体ID
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 字段
	Fields []DBField `json:"fieldList" gorm:"column:-"`
	// 备注信息
	Comment string `json:"comment" gorm:"column:comment"`
}

func (receiver DBStructInfo) TableName() string {
	return "struct_info"
}
