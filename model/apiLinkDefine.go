package model

// LinkDefineInfo 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type LinkDefineInfo struct {
	DBLinkDefine
	HasInstance bool              `json:"hasInstance"`
	ParamList   []LinkParamDefine `json:"paramList"`
}

// GetLinkDefineResp 获取 project define 信息返回结构
type GetLinkDefineResp struct {
	BaseResp
	// 结构体信息
	Info LinkDefineInfo `json:"info"`
}

// CreateLinkDefineReq 新建 project define 信息请求结构
type CreateLinkDefineReq struct {
	// 项目名
	Name string `json:"name"`
	// 字段注释
	Comment string `json:"comment"`
	// 参数列表
	ParamList []LinkParamDefine `json:"paramList"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateLinkDefineReq) ToDBStruct() (DBLinkDefine, []DBLinkParamDefine) {
	var result = DBLinkDefine{
		Name:    receiver.Name,
		Comment: receiver.Comment,
	}
	list := make([]DBLinkParamDefine, len(receiver.ParamList))

	for i, p := range receiver.ParamList {
		list[i] = DBLinkParamDefine{
			Location: p.Location,
			PType:    p.PType,
			Name:     p.Name,
			Comment:  p.Comment,
			IsList:   p.IsList,
			Sort:     p.Sort,
		}
	}

	return result, list
}

// CreateLinkDefineResp 新建 project define 信息返回结构
type CreateLinkDefineResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateLinkDefineReq 更新 project define 信息请求结构
type UpdateLinkDefineReq struct {
	LinkDefineInfo
}

// UpdateLinkDefineResp 更新 project define 信息返回结构
type UpdateLinkDefineResp struct {
	BaseResp
	// 更新的ID
	ID int `json:"id"`
}

// DeleteLinkDefineResp 删除 project define 信息返回结构
type DeleteLinkDefineResp struct {
	BaseResp
	// 删除的ID
	ID int `json:"id"`
}

// GetLinkDefineListResp 获取 project define 列表返回结构
type GetLinkDefineListResp struct {
	BaseResp
	// 项目列表
	List []LinkDefineInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// LinkParamDefine link param define 信息
type LinkParamDefine struct {
	DBLinkParamDefine
	HasInstance bool `json:"hasInstance"`
}
