package service

import "AmazingSaltedFish/model"

// GetNodeInstance 获取 node instance 信息
func GetNodeInstance(id int) (model.GetNodeInstanceResp, error) {
	var result model.GetNodeInstanceResp

	return result, nil
}

// GetNodeInstanceList 获取 node instance 列表信息
func GetNodeInstanceList(reqInfo model.PageReq) (model.GetNodeInstanceListResp, error) {
	var result model.GetNodeInstanceListResp

	return result, nil
}

// CreateNodeInstance 新建 node instance 信息
func CreateNodeInstance(reqInfo model.CreateNodeInstanceReq) (model.CreateNodeInstanceResp, error) {
	var result model.CreateNodeInstanceResp

	return result, nil
}

// UpdateNodeInstance 更新 node instance 信息
func UpdateNodeInstance(reqInfo model.UpdateNodeInstanceReq) (model.UpdateNodeInstanceResp, error) {
	var result model.UpdateNodeInstanceResp

	return result, nil
}

// DeleteNodeInstance 删除 node instance 信息
func DeleteNodeInstance(id int) (model.DeleteNodeInstanceResp, error) {
	var result model.DeleteNodeInstanceResp

	return result, nil
}
