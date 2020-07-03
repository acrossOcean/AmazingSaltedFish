package service

import "AmazingSaltedFish/model"

// GetProjectInstance 获取 Project instance 信息
func GetProjectInstance(id int) (model.GetProjectInstanceResp, error) {
	var result model.GetProjectInstanceResp

	return result, nil
}

// GetProjectInstanceList 获取 Project instance 列表信息
func GetProjectInstanceList(reqInfo model.PageReq) (model.GetProjectInstanceListResp, error) {
	var result model.GetProjectInstanceListResp

	return result, nil
}

// CreateProjectInstance 新建 Project instance 信息
func CreateProjectInstance(reqInfo model.CreateProjectInstanceReq) (model.CreateProjectInstanceResp, error) {
	var result model.CreateProjectInstanceResp

	return result, nil
}

// UpdateProjectInstance 更新 Project instance 信息
func UpdateProjectInstance(reqInfo model.UpdateProjectInstanceReq) (model.UpdateProjectInstanceResp, error) {
	var result model.UpdateProjectInstanceResp

	return result, nil
}

// DeleteProjectInstance 删除 Project instance 信息
func DeleteProjectInstance(id int) (model.DeleteProjectInstanceResp, error) {
	var result model.DeleteProjectInstanceResp

	return result, nil
}
