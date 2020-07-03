package model

import "AmazingSaltedFish/utils/check"

// 结构信息
// 包含结构定义, 如果含有实现, 也包含实现信息
type NodeDefineInfo struct {
	DBNodeDefine
	HasInstance bool              `json:"hasInstance"`
	ParamList   []NodeParamDefine `json:"paramList"`
}

type GetNodeDefineResp struct {
	BaseResp
	// node信息
	Info NodeDefineInfo `json:"info"`
}

type GetNodeDefineListResp struct {
	BaseResp
	// node列表
	List []NodeDefineInfo `json:"list"`
	// 总数据条数
	Sum int `json:"sum"`
}

type CreateNodeDefineReq struct {
	// 所属link id
	BelongLinkDefineId int `json:"belongLinkDefineId"`
	// node名
	Name string `json:"name"`
	// 注释
	Comment string `json:"comment"`
	// 前一个nodeID , 第一个写-1
	PreNodeDefineId int `json:"preNodeDefineId"`
	// 参数列表
	ParamList []NodeParamDefine `json:"paramList"`
}

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

type CreateNodeDefineResp struct {
	BaseResp
	// 创建成功后的ID
	Id int `json:"id"`
}

type UpdateNodeDefineReq struct {
	NodeDefineInfo
}

func (receiver UpdateNodeDefineReq) Check() bool {
	if !check.PassCheck(
		check.NewIDChecker(receiver.Id),
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

type UpdateNodeDefineResp struct {
	BaseResp
	// 更新的ID
	Id int `json:"id"`
}

type DeleteNodeDefineResp struct {
	BaseResp
	// 删除的ID
	Id int `json:"id"`
}

// 项目参数信息
type NodeParamDefine struct {
	DBNodeParamDefine
	HasInstance bool `json:"hasInstance"`
}

func (receiver NodeParamDefine) ToDBStruct() DBNodeParamDefine {
	var result = DBNodeParamDefine{
		Id:              receiver.Id,
		NodeDefineId:    receiver.NodeDefineId,
		Location:        receiver.Location,
		PType:           receiver.PType,
		Name:            receiver.Name,
		Comment:         receiver.Comment,
		IsList:          receiver.IsList,
		Sort:            receiver.Sort,
		StructId:        receiver.StructId,
		MapKeyParamId:   receiver.MapKeyParamId,
		MapValueParamId: receiver.MapValueParamId,
	}

	return result
}
