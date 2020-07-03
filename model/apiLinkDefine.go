package model

// 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type LinkDefineInfo struct {
	DBLinkDefine
	HasInstance bool              `json:"hasInstance"`
	ParamList   []LinkParamDefine `json:"paramList"`
}

type GetLinkDefineResp struct {
	BaseResp
	// 结构体信息
	Info LinkDefineInfo `json:"info"`
}

type CreateLinkDefineReq struct {
	// 项目名
	Name string `json:"name"`
	// 字段注释
	Comment string `json:"comment"`
	// 参数列表
	ParamList []LinkParamDefine `json:"paramList"`
}

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

type CreateLinkDefineResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateLinkDefineReq struct {
	LinkDefineInfo
}

type UpdateLinkDefineResp struct {
	BaseResp
	// 更新的ID
	Id int `json:"id"`
}

type DeleteLinkDefineResp struct {
	BaseResp
	// 删除的ID
	Id int `json:"id"`
}

type GetLinkDefineListResp struct {
	BaseResp
	// 项目列表
	List []LinkDefineInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// 项目参数信息
type LinkParamDefine struct {
	DBLinkParamDefine
	HasInstance bool `json:"hasInstance"`
}
