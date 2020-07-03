package model

// NodeInstanceInfo 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type NodeInstanceInfo struct {
	DBNodeInstance
	ParamList []NodeParamInstance `json:"paramList"`
}

// GetNodeInstanceResp 获取 node instance 信息返回结构
type GetNodeInstanceResp struct {
	BaseResp
	// 结构体信息
	Info NodeInstanceInfo `json:"info"`
}

// CreateNodeInstanceReq 新建 node instance 信息请求结构
type CreateNodeInstanceReq struct {
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
	ParamList []NodeParamInstance `json:"paramList"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateNodeInstanceReq) ToDBStruct() (DBNodeInstance, []DBNodeParamInstance) {
	var result = DBNodeInstance{
		DefineID: receiver.DefineID,
		Name:     receiver.Name,
		Comment:  receiver.Comment,
		//GenerateLanguage:    receiver.GenerateLanguage,
		//FirstLinkDefineID:   receiver.FirstLinkDefineID,
		//FirstLinkInstanceID: receiver.FirstLinkInstanceID,
	}

	list := make([]DBNodeParamInstance, len(receiver.ParamList))

	for i, p := range receiver.ParamList {
		list[i] = DBNodeParamInstance{
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

// CreateNodeInstanceResp 新建 node instance 信息返回结构
type CreateNodeInstanceResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateNodeInstanceReq 更新 node instance 信息请求结构
type UpdateNodeInstanceReq struct {
	NodeInstanceInfo
}

// UpdateNodeInstanceResp 更新 node instance 信息返回结构
type UpdateNodeInstanceResp struct {
	BaseResp
	// 更新的ID
	ID int `json:"id"`
}

// DeleteNodeInstanceResp 删除 node instance 信息返回结构
type DeleteNodeInstanceResp struct {
	BaseResp
	// 删除的ID
	ID int `json:"id"`
}

// GetNodeInstanceListResp 获取 node instance 列表返回结构
type GetNodeInstanceListResp struct {
	BaseResp
	// 项目列表
	List []NodeInstanceInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// NodeParamInstance node param instance 信息
type NodeParamInstance struct {
	DBNodeParamInstance
}
