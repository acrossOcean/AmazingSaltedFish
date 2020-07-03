package model

// 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type LinkInstanceInfo struct {
	DBLinkInstance
	ParamList []LinkParamInstance `json:"paramList"`
}

type GetLinkInstanceResp struct {
	BaseResp
	// 结构体信息
	Info LinkInstanceInfo `json:"info"`
}

type CreateLinkInstanceReq struct {
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
	ParamList []LinkParamInstance `json:"paramList"`
}

func (receiver CreateLinkInstanceReq) ToDBStruct() (DBLinkInstance, []DBLinkParamInstance) {
	var result = DBLinkInstance{
		DefineId: receiver.DefineId,
		Name:     receiver.Name,
		Comment:  receiver.Comment,
		//GenerateLanguage:    receiver.GenerateLanguage,
		//FirstLinkDefineId:   receiver.FirstLinkDefineId,
		//FirstLinkInstanceId: receiver.FirstLinkInstanceId,
	}

	list := make([]DBLinkParamInstance, len(receiver.ParamList))

	for i, p := range receiver.ParamList {
		list[i] = DBLinkParamInstance{
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

type CreateLinkInstanceResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateLinkInstanceReq struct {
	LinkInstanceInfo
}

type UpdateLinkInstanceResp struct {
	BaseResp
	// 更新的ID
	Id int `json:"id"`
}

type DeleteLinkInstanceResp struct {
	BaseResp
	// 删除的ID
	Id int `json:"id"`
}

type GetLinkInstanceListResp struct {
	BaseResp
	// 项目列表
	List []LinkInstanceInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// 项目参数信息
type LinkParamInstance struct {
	DBLinkParamInstance
}
