package model

// ProjectInstanceInfo  project instance 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type ProjectInstanceInfo struct {
	DBProjectInstance
	ParamList []ProjectParamInstance `json:"paramList"`
}

// GetProjectInstanceResp 获取 project instance 信息返回结构
type GetProjectInstanceResp struct {
	BaseResp
	// 结构体信息
	Info ProjectInstanceInfo `json:"info"`
}

// CreateProjectInstanceReq 新建 project instance 信息请求结构
type CreateProjectInstanceReq struct {
	// 对应 定义 id
	DefineID int `json:"defineId"`
	// 工程名
	Name string `json:"name"`
	// 注释
	Comment string `json:"comment"`
	// 实现工程所用语言及版本
	GenerateLanguage string `json:"generateLanguage"`

	// 第一个 link 信息
	FirstLinkDefineID   int `json:"firstLinkDefineId"`
	FirstLinkInstanceID int `json:"firstLinkInstanceId"`

	// 参数信息
	ParamList []ProjectParamInstance `json:"paramList"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateProjectInstanceReq) ToDBStruct() (DBProjectInstance, []DBProjectParamInstance) {
	var result = DBProjectInstance{
		DefineID:            receiver.DefineID,
		Name:                receiver.Name,
		Comment:             receiver.Comment,
		GenerateLanguage:    receiver.GenerateLanguage,
		FirstLinkDefineID:   receiver.FirstLinkDefineID,
		FirstLinkInstanceID: receiver.FirstLinkInstanceID,
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
			InputVarDefineID:   p.InputVarDefineID,
			InputVarInstanceID: p.InputVarInstanceID,
			OutputType:         p.OutputType,
			OutputConst:        p.OutputConst,
		}
	}

	return result, list
}

// CreateProjectInstanceResp 新建 project instance 信息返回结构
type CreateProjectInstanceResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateProjectInstanceReq 更新 project instance 信息请求结构
type UpdateProjectInstanceReq struct {
	ProjectInstanceInfo
}

// UpdateProjectInstanceResp 更新 project instance 信息返回结构
type UpdateProjectInstanceResp struct {
	BaseResp
	// 更新的ID
	ID int `json:"id"`
}

// DeleteProjectInstanceResp 删除 project instance 信息返回结构
type DeleteProjectInstanceResp struct {
	BaseResp
	// 删除的ID
	ID int `json:"id"`
}

// GetProjectInstanceListResp 获取 project instance 列表返回结构
type GetProjectInstanceListResp struct {
	BaseResp
	// 项目列表
	List []ProjectInstanceInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// ProjectParamInstance project param instance 信息
type ProjectParamInstance struct {
	DBProjectParamInstance
}
