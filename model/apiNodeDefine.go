package model

import "AmazingSaltedFish/utils/check"

// NodeDefineInfo 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type NodeDefineInfo struct {
	DBNodeDefine
	HasInstance bool              `json:"hasInstance"`
	ParamList   []NodeParamDefine `json:"paramList"`
}

// GetNodeDefineResp 获取 project define 信息返回结构
type GetNodeDefineResp struct {
	BaseResp
	// node信息
	Info NodeDefineInfo `json:"info"`
}

// CreateNodeDefineReq 新建 project define 信息请求结构
type CreateNodeDefineReq struct {
	// 所属link id
	BelongLinkDefineID int `json:"belongLinkDefineId"`
	// node名
	Name string `json:"name"`
	// 注释
	Comment string `json:"comment"`
	// 前一个nodeID , 第一个写-1
	PreNodeDefineID int `json:"preNodeDefineId"`
	// 参数列表
	ParamList []NodeParamDefine `json:"paramList"`
}

// Check 检查请求参数是否合法
func (receiver CreateNodeDefineReq) Check() bool {
	if !check.PassCheck(
		check.NewStrChecker(receiver.Name, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
		check.NewStrChecker(receiver.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
		check.NewStrChecker(receiver.Name, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
		check.NewStrChecker(receiver.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
	) {
		return false
	}

	for _, param := range receiver.ParamList {
		if !check.PassCheck(
			check.NewIntChecker(param.PType.ToInt(),
				check.NewIntCheckOptionIN(GetAllFieldTypeInt()),
			),
			check.NewStrChecker(param.Name, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
			check.NewStrChecker(param.Name, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
			check.NewStrChecker(param.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
			check.NewStrChecker(param.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
		) {
			return false
		}
	}

	return true
}

// ToDBStruct 转换为数据库对应结构
func (receiver CreateNodeDefineReq) ToDBStruct() (DBNodeDefine, []DBNodeParamDefine) {
	var result = DBNodeDefine{
		Name:    receiver.Name,
		Comment: receiver.Comment,
	}
	list := make([]DBNodeParamDefine, len(receiver.ParamList))

	for i, p := range receiver.ParamList {
		list[i] = p.ToDBStruct()
	}

	return result, list
}

// CreateNodeDefineResp 新建 project define 信息返回结构
type CreateNodeDefineResp struct {
	BaseResp
	// 创建成功后的ID
	ID int `json:"id"`
}

// UpdateNodeDefineReq 更新 project define 信息请求结构
type UpdateNodeDefineReq struct {
	NodeDefineInfo
}

// Check 检查请求参数是否合法
func (receiver UpdateNodeDefineReq) Check() bool {
	if !check.PassCheck(
		check.NewIDChecker(receiver.ID),
		check.NewStrChecker(receiver.Name, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
		check.NewStrChecker(receiver.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
		check.NewStrChecker(receiver.Name, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
		check.NewStrChecker(receiver.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
	) {
		return false
	}

	for _, param := range receiver.ParamList {
		if !check.PassCheck(
			check.NewIntChecker(param.PType.ToInt(),
				check.NewIntCheckOptionIN(GetAllFieldTypeInt()),
			),
			check.NewStrChecker(param.Name, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
			check.NewStrChecker(param.Name, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
			check.NewStrChecker(param.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenGT, 0)),
			check.NewStrChecker(param.Comment, check.NewStrCheckOptionInt(check.StrOperatorLenLT, 200)),
		) {
			return false
		}
	}

	return true
}

// UpdateNodeDefineResp 更新 project define 信息返回结构
type UpdateNodeDefineResp struct {
	BaseResp
	// 更新的ID
	ID int `json:"id"`
}

// DeleteNodeDefineResp 删除 project define 信息返回结构
type DeleteNodeDefineResp struct {
	BaseResp
	// 删除的ID
	ID int `json:"id"`
}

// GetNodeDefineListResp  获取 project define 列表返回结构
type GetNodeDefineListResp struct {
	BaseResp
	// node列表
	List []NodeDefineInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

// NodeParamDefine node param define 信息
type NodeParamDefine struct {
	DBNodeParamDefine
	HasInstance bool `json:"hasInstance"`
}

// ToDBStruct 转换为数据库对应结构
func (receiver NodeParamDefine) ToDBStruct() DBNodeParamDefine {
	var result = DBNodeParamDefine{
		ID:              receiver.ID,
		NodeDefineID:    receiver.NodeDefineID,
		Location:        receiver.Location,
		PType:           receiver.PType,
		Name:            receiver.Name,
		Comment:         receiver.Comment,
		IsList:          receiver.IsList,
		Sort:            receiver.Sort,
		StructID:        receiver.StructID,
		MapKeyParamID:   receiver.MapKeyParamID,
		MapValueParamID: receiver.MapValueParamID,
	}

	return result
}
