package service

import (
	"AmazingSaltedFish/model"
	"AmazingSaltedFish/utils"

	"github.com/acrossOcean/log"
	"github.com/jinzhu/gorm"
)

// GetNodeDefine 获取 node define 信息
func GetNodeDefine(id int) (model.GetNodeDefineResp, error) {
	var result model.GetNodeDefineResp
	var info model.DBNodeDefine

	err := GetDB().Where("id = ?", id).First(&info).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Error("根据 id 获取 node 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeResourceNotExist, err)
	} else if err != nil {
		log.Error("根据 id 获取 node 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	var count int
	err = GetDB().Table(model.DBNodeInstance{}.TableName()).Where("define_id = ?", info.Id).Count(&count).Error
	if err != nil {
		log.Error("根据 node define ID(%v) 获取 对应 实现 信息数量 错误:%s", info.ID, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	params, err := GetNodeParamsDefineByNodeID(info.Id)
	if err != nil {
		log.Error("根据 node define id 获取参数信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	result.Info.ParamList = params
	result.Info.DBNodeDefine = info
	result.Info.HasInstance = count > 0
	result.SetSuccess()

	return result, nil
}

// GetNodeDefineList 获取 node define 列表信息
func GetNodeDefineList(reqInfo model.PageReq) (model.GetNodeDefineListResp, error) {
	var result model.GetNodeDefineListResp
	var list = make([]model.NodeDefineInfo, 0)
	var infos = make([]model.DBNodeDefine, 0)
	var count int
	if err := GetDB().Table(model.DBNodeDefine{}.TableName()).Count(&count).Error; err != nil {
		log.Error("获取 node define表 数据总数错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := GetDB().Order("id DESC").
		Limit(reqInfo.PageSize).Offset((reqInfo.PageNum - 1) * reqInfo.PageSize).
		Find(&infos).Error; err != nil {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	ids := make([]int, len(infos))
	for i, info := range infos {
		ids[i] = info.Id
	}

	// 查询
	var instanceCountList = make([]struct {
		ID    int `gorm:"column:id"`
		Count int `gorm:"column:count"`
	}, 0)

	if err := GetDB().Table(model.DBNodeInstance{}.TableName()).
		Select("id as id, count(1) as count").
		Where("define_id in (?)", ids).
		Group("define_id").Scan(&instanceCountList).Error; err != nil {
		log.Error("根据 node define id 获取 实现数量错误:", err.Error())
	}
	instanceMap := make(map[int]int)
	for _, count := range instanceCountList {
		instanceMap[count.ID] = count.Count
	}

	paramMap, err := GetNodeParamsDefineByNodeIds(ids)
	if err != nil {
		log.Error("根据 node define ids 获取 参数信息 错误:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, info := range infos {
		var temp model.NodeDefineInfo
		temp.HasInstance = instanceMap[info.Id] > 0
		temp.ParamList = paramMap[info.Id]
		temp.DBNodeDefine = info

		list = append(list, temp)
	}

	result.Sum = count
	result.List = list
	result.SetSuccess()

	return result, nil
}

// CreateNodeDefine 新建 node define 信息
func CreateNodeDefine(reqInfo model.CreateNodeDefineReq) (model.CreateNodeDefineResp, error) {
	var result model.CreateNodeDefineResp

	tx := GetDB().Begin()

	// 新建一个
	info, paramList := reqInfo.ToDBStruct()

	if err := tx.Save(&info).Error; err != nil {
		tx.Rollback()
		log.Error("保存 node define 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, param := range paramList {
		param.NodeDefineID = info.Id
		if err := tx.Save(&param).Error; err != nil {
			tx.Rollback()
			log.Error("保存 node param define 信息错误:%s", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	tx.Commit()
	result.ID = info.Id
	result.SetSuccess()

	return result, nil
}

// UpdateNodeDefine 更新 node define 信息
func UpdateNodeDefine(reqInfo model.UpdateNodeDefineReq) (model.UpdateNodeDefineResp, error) {
	var result model.UpdateNodeDefineResp

	// 先获取老数据
	var info model.DBNodeDefine
	if err := GetDB().Where("id = ?", reqInfo.Id).First(&info).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Error("要更新的 node define 数据(%d)不存在", reqInfo.Id)
			return result, model.ReturnWithCode(model.CodeResourceNotExist, err)
		}

		log.Error("获取要更新的 node define 数据(%d)错误:%s", reqInfo.ID, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	// 暂不考虑同步更新实现的问题, 先只更新定义
	tx := GetDB().Begin()
	if err := tx.Save(&reqInfo.DBNodeDefine).Error; err != nil {
		tx.Rollback()
		log.Error("更新 node define 错误:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, param := range reqInfo.ParamList {
		param.NodeDefineID = reqInfo.Id
		if err := tx.Save(&param.DBNodeParamDefine).Error; err != nil {
			tx.Rollback()
			log.Error("更新 node param define 错误:", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	// 如果更改了node 顺序,那么前后节点都重新更改下
	if reqInfo.PreNodeDefineID != info.PreNodeDefineID {
		if err := tx.Where("pre_node_define_id = ?", info.Id).Update("pre_node_define_id", info.PreNodeDefineId).Error; err != nil {
			tx.Rollback()
			log.Error("更新 node define 时, 更新 node define 顺序错误:", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}

		if err := tx.Where("pre_node_define_id = ?", reqInfo.PreNodeDefineId).Update("pre_node_define_id", info.Id).Error; err != nil {
			tx.Rollback()
			log.Error("更新 node define 时, 更新 node define 顺序错误:", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	tx.Commit()
	result.ID = reqInfo.Id
	result.SetSuccess()

	return result, nil
}

// DeleteNodeDefine 删除 node define 信息
func DeleteNodeDefine(id int) (model.DeleteNodeDefineResp, error) {
	var result model.DeleteNodeDefineResp

	// 先获取老的定义
	var info model.DBNodeDefine
	if err := GetDB().Where("id = ?", id).First(&info).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Error("要删除的资源(%d)不存在", id)
			return result, model.ReturnWithCode(model.CodeResourceNotExist, err)
		}

		log.Error("获取要删除的资源(%d)错误:%s", id, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	// 删除, 同时删除定义和实例
	tx := GetDB().Begin()

	if err := tx.Where("id = ?", id).Delete(&model.DBNodeDefine{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 node define 失败:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := tx.Where("node_define_id = ?", id).Delete(&model.DBNodeParamDefine{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 node define param 失败:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := tx.Where("define_id = ?", id).Delete(&model.DBNodeInstance{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 node instance 失败:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := tx.Where("node_define_id = ?", id).Delete(&model.DBNodeParamInstance{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 node instance param 失败:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	// 将当前节点下一个节点的记录更新为当前节点的
	if err := tx.Where("pre_node_define_id = ?", id).Update("pre_node_define_id", info.PreNodeDefineId).Error; err != nil {
		tx.Rollback()
		log.Error("更新 node define 失败:", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	tx.Commit()
	result.ID = id
	result.SetSuccess()

	return result, nil
}

// GetNodeParamsDefineByNodeID 根据 node define id 获取 定义的参数信息
func GetNodeParamsDefineByNodeID(id int) ([]model.NodeParamDefine, error) {
	var result = make([]model.NodeParamDefine, 0)

	// 如果传入id 为0, 那么返回空串
	if id == 0 {
		return result, nil
	}

	var list = make([]model.DBNodeParamDefine, 0)
	if err := GetDB().Where("node_define_id = ?", id).Find(&list).Error; err != nil {
		log.Error("根据node define ID 获取 params 信息错误:", err.Error())
		return result, err
	}

	paramDefineIds := make([]int, len(list))
	paramMap := make(map[int]model.DBNodeParamDefine)
	for i, param := range list {
		paramDefineIds[i] = param.Id
		paramMap[param.Id] = param
	}

	// 查询
	var instanceCountList = make([]struct {
		ID    int `gorm:"column:id"`
		Count int `gorm:"column:count"`
	}, 0)

	if err := GetDB().Table(model.DBNodeInstance{}.TableName()).
		Select("id as id, count(1) as count").
		Where("define_id in (?)", paramDefineIds).
		Group("define_id").Scan(&instanceCountList).Error; err != nil {
		log.Error("根据 node param define id 获取 实现数量错误:", err.Error())
	}

	for _, count := range instanceCountList {
		var temp model.NodeParamDefine

		temp.HasInstance = count.Count > 0
		temp.DBNodeParamDefine = paramMap[count.ID]

		result = append(result, temp)
	}

	return result, nil
}

// GetNodeParamsDefineByNodeIds 根据 node define ids 获取 定义的参数信息
func GetNodeParamsDefineByNodeIds(ids []int) (map[int][]model.NodeParamDefine, error) {
	var result = make(map[int][]model.NodeParamDefine)

	ids = utils.IntList(utils.IntList(ids).RemoveDuplicate()).RemoveZero()

	// 如果传入id 为0, 那么返回空串
	if len(ids) == 0 {
		return result, nil
	}

	var list = make([]model.DBNodeParamDefine, 0)
	if err := GetDB().Where("node_define_id in (?)", ids).Find(&list).Error; err != nil {
		log.Error("根据node define ID 获取 params 信息错误:", err.Error())
		return result, err
	}

	paramDefineIds := make([]int, len(list))
	paramMap := make(map[int]model.DBNodeParamDefine)
	for i, param := range list {
		paramDefineIds[i] = param.Id
		paramMap[param.Id] = param
	}

	// 查询
	var instanceCountList = make([]struct {
		ID    int `gorm:"column:id"`
		Count int `gorm:"column:count"`
	}, 0)

	if err := GetDB().Table(model.DBNodeParamInstance{}.TableName()).
		Select("id as id, count(1) as count").
		Where("define_id in (?)", paramDefineIds).
		Group("define_id").Scan(&instanceCountList).Error; err != nil {
		log.Error("根据 node param define id 获取 实现数量错误:", err.Error())
	}

	for _, count := range instanceCountList {
		var temp model.NodeParamDefine

		temp.HasInstance = count.Count > 0
		temp.DBNodeParamDefine = paramMap[count.ID]

		list := result[paramMap[count.ID].NodeDefineId]
		list = append(list, temp)
		result[paramMap[count.ID].NodeDefineId] = list
	}

	return result, nil
}
