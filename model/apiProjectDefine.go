package model

// 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type ProjectDefineInfo struct {
	DBProjectDefine
	HasInstance bool                 `json:"hasInstance"`
	ParamList   []ProjectParamDefine `json:"paramList"`
}

type GetProjectDefineResp struct {
	BaseResp
	// 结构体信息
	Info ProjectDefineInfo `json:"info"`
}

type CreateProjectDefineReq struct {
	// 项目名
	Name string `json:"name"`
	// 字段注释
	Comment string `json:"comment"`
	// 参数列表
	ParamList []ProjectParamDefine `json:"paramList"`
}

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

type CreateProjectDefineResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateProjectDefineReq struct {
	ProjectDefineInfo
}

type UpdateProjectDefineResp struct {
	BaseResp
	// 更新的ID
	Id int `json:"id"`
}

type DeleteProjectDefineResp struct {
	BaseResp
	// 删除的ID
	Id int `json:"id"`
}

type GetProjectDefineListResp struct {
	BaseResp
	// 项目列表
	List []ProjectDefineInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// 项目参数信息
type ProjectParamDefine struct {
	DBProjectParamDefine
	HasInstance bool `json:"hasInstance"`
}
