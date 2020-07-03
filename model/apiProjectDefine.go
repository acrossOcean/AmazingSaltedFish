package model

// ProjectDefineInfo 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type ProjectDefineInfo struct {
	DBProjectDefine
	HasInstance bool                 `json:"hasInstance"`
	ParamList   []ProjectParamDefine `json:"paramList"`
}

// GetProjectDefineResp 获取 project define 信息返回结构
type GetProjectDefineResp struct {
	BaseResp
	// 结构体信息
	Info ProjectDefineInfo `json:"info"`
}

// CreateProjectDefineReq 新建 project define 信息请求结构
type CreateProjectDefineReq struct {
	// 项目名
	Name string `json:"name"`
	// 字段注释
	Comment string `json:"comment"`
	// 参数列表
	ParamList []ProjectParamDefine `json:"paramList"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateProjectDefineReq) ToDBStruct() (DBProjectDefine, []DBProjectParamDefine) {
	var result = DBProjectDefine{
		Name:    receiver.Name,
		Comment: receiver.Comment,
	}
	list := make([]DBProjectParamDefine, len(receiver.ParamList))

	for i, p := range receiver.ParamList {
		list[i] = DBProjectParamDefine{
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

// CreateProjectDefineResp 新建 project define 信息返回结构
type CreateProjectDefineResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateProjectDefineReq 更新 project define 信息请求结构
type UpdateProjectDefineReq struct {
	ProjectDefineInfo
}

// UpdateProjectDefineResp 更新 project define 信息返回结构
type UpdateProjectDefineResp struct {
	BaseResp
	// 更新的ID
	ID int `json:"id"`
}

// DeleteProjectDefineResp 删除 project define 信息返回结构
type DeleteProjectDefineResp struct {
	BaseResp
	// 删除的ID
	ID int `json:"id"`
}

// GetProjectDefineListResp 获取 project define 列表返回结构
type GetProjectDefineListResp struct {
	BaseResp
	// 项目列表
	List []ProjectDefineInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// ProjectParamDefine project param define 信息
type ProjectParamDefine struct {
	DBProjectParamDefine
	HasInstance bool `json:"hasInstance"`
}
