package model

// LinkInstanceInfo 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type LinkInstanceInfo struct {
	DBLinkInstance
	ParamList []LinkParamInstance `json:"paramList"`
}

// GetLinkInstanceResp 获取 link instance 信息返回结构
type GetLinkInstanceResp struct {
	BaseResp
	// 结构体信息
	Info LinkInstanceInfo `json:"info"`
}

// CreateLinkInstanceReq 新建 link instance 信息请求结构
type CreateLinkInstanceReq struct {
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
	ParamList []LinkParamInstance `json:"paramList"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateLinkInstanceReq) ToDBStruct() (DBLinkInstance, []DBLinkParamInstance) {
	var result = DBLinkInstance{
		DefineID: receiver.DefineID,
		Name:     receiver.Name,
		Comment:  receiver.Comment,
		//GenerateLanguage:    receiver.GenerateLanguage,
		//FirstLinkDefineID:   receiver.FirstLinkDefineID,
		//FirstLinkInstanceID: receiver.FirstLinkInstanceID,
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
			InputVarDefineID:   p.InputVarDefineID,
			InputVarInstanceID: p.InputVarInstanceID,
			OutputType:         p.OutputType,
			OutputConst:        p.OutputConst,
		}
	}

	return result, list
}

// CreateLinkInstanceResp 新建 link instance 信息返回结构
type CreateLinkInstanceResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateLinkInstanceReq 更新 link instance 信息请求结构
type UpdateLinkInstanceReq struct {
	LinkInstanceInfo
}

// UpdateLinkInstanceResp 更新 link instance 信息返回结构
type UpdateLinkInstanceResp struct {
	BaseResp
	// 更新的ID
	ID int `json:"id"`
}

// DeleteLinkInstanceResp 删除 link instance 信息返回结构
type DeleteLinkInstanceResp struct {
	BaseResp
	// 删除的ID
	ID int `json:"id"`
}

// GetLinkInstanceListResp 获取 link instance 列表返回结构
type GetLinkInstanceListResp struct {
	BaseResp
	// 项目列表
	List []LinkInstanceInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// LinkParamInstance link param instance 信息
type LinkParamInstance struct {
	DBLinkParamInstance
}
