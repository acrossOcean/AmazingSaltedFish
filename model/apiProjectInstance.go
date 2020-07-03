package model

// 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type ProjectInstanceInfo struct {
	DBProjectInstance
	ParamList []ProjectParamInstance `json:"paramList"`
}

type GetProjectInstanceResp struct {
	BaseResp
	// 结构体信息
	Info ProjectInstanceInfo `json:"info"`
}

type CreateProjectInstanceReq struct {
	// 对应 定义 id
	DefineId int `json:"defineId"`
	// 工程名
	Name string `json:"name"`
	// 注释
	Comment string `json:"comment"`
	// 实现工程所用语言及版本
	GenerateLanguage string `json:"generateLanguage"`

	// 第一个 link 信息
	FirstLinkDefineId   int `json:"firstLinkDefineId"`
	FirstLinkInstanceId int `json:"firstLinkInstanceId"`

	// 参数信息
	ParamList []ProjectParamInstance `json:"paramList"`
}

func (receiver CreateProjectInstanceReq) ToDBStruct() (DBProjectInstance, []DBProjectParamInstance) {
	var result = DBProjectInstance{
		DefineId:            receiver.DefineId,
		Name:                receiver.Name,
		Comment:             receiver.Comment,
		GenerateLanguage:    receiver.GenerateLanguage,
		FirstLinkDefineId:   receiver.FirstLinkDefineId,
		FirstLinkInstanceId: receiver.FirstLinkInstanceId,
	}

	list := make([]DBProjectParamInstance, len(receiver.ParamList))

	for i, p := range receiver.ParamList {
		list[i] = DBProjectParamInstance{
			Location:           p.Location,
			PType:              p.PType,
			Name:               p.Name,
			Comment:            p.Comment,
			InputType:          p.InputType,
			InputConst:         p.InputConst,
			InputVarIsLink:     p.InputVarIsLink,
			InputVarDefineId:   p.InputVarDefineId,
			InputVarInstanceId: p.InputVarInstanceId,
			OutputType:         p.OutputType,
			OutputConst:        p.OutputConst,
		}
	}

	return result, list
}

type CreateProjectInstanceResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateProjectInstanceReq struct {
	ProjectInstanceInfo
}

type UpdateProjectInstanceResp struct {
	BaseResp
	// 更新的ID
	Id int `json:"id"`
}

type DeleteProjectInstanceResp struct {
	BaseResp
	// 删除的ID
	Id int `json:"id"`
}

type GetProjectInstanceListResp struct {
	BaseResp
	// 项目列表
	List []ProjectInstanceInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// 项目参数信息
type ProjectParamInstance struct {
	DBProjectParamInstance
}
