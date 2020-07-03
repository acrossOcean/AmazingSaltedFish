package service

import (
	"AmazingSaltedFish/model"
	"sort"

	"github.com/jinzhu/gorm"

	"github.com/acrossOcean/log"
)

// GetStruct 获取结构体详情
func GetStruct(id int) (model.GetStructResp, error) {
	var result model.GetStructResp
	var info model.DBStructInfo

	err := GetDB().Where("id = ?", id).First(&info).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeResourceNotExist, err)
	} else if err != nil {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	fields := make([]model.DBField, 0)
	err = GetDB().Where("parent_id = ?", info.ID).Find(&fields).Error
	if err != nil {
		log.Error("根据 结构体ID(%v) 获取 对应字段信息 错误:%s", info.ID, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}
	sort.Sort(model.FieldList(fields))

	info.Fields = fields
	result.Info = info
	result.SetSuccess()

	return result, nil
}

// GetStructList 获取结构体列表
func GetStructList(req model.PageReq) (model.GetStructListResp, error) {
	var result model.GetStructListResp
	var infos []model.DBStructInfo
	var count int

	if err := GetDB().Table(model.DBStructInfo{}.TableName()).Count(&count).Error; err != nil {
		log.Error("获取 结构体表 数据总数错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := GetDB().Order("id DESC").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&infos).Error; err != nil {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	structIds := make([]int, len(infos))
	for i, info := range infos {
		structIds[i] = info.ID
	}

	fields := make([]model.DBField, 0)
	if err := GetDB().Where("parent_id IN (?)", structIds).Find(&fields).Error; err != nil {
		log.Error("根据 结构体IDs(%v) 获取 对应字段信息 错误:%s", structIds, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	fieldMap := make(map[int][]model.DBField)
	for _, f := range fields {
		if _, ok := fieldMap[f.ParentId]; !ok {
			fieldMap[f.ParentId] = make([]model.DBField, 0)
		}

		fieldMap[f.ParentId] = append(fieldMap[f.ParentId], f)
	}

	for i, info := range infos {
		if list, ok := fieldMap[info.ID]; ok {
			sort.Sort(model.FieldList(list))
			info.Fields = list
			infos[i] = info
		} else {
			info.Fields = make([]model.DBField, 0)
			infos[i] = info
		}
	}

	result.Sum = count
	result.List = infos
	result.SetSuccess()

	return result, nil
}

// CreateStruct 新建结构体
func CreateStruct(reqInfo model.CreateStructReq) (model.CreateStructResp, error) {
	var result model.CreateStructResp

	tx := GetDB().Begin()
	data := reqInfo.ToDBStruct()

	if err := tx.Save(&data).Error; err != nil {
		tx.Rollback()
		log.Error("保存 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, f := range data.Fields {
		f.ParentID = data.ID
		if err := tx.Save(&f).Error; err != nil {
			tx.Rollback()
			log.Error("保存 结构体的字段 信息错误:%s", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	tx.Commit()
	result.ID = data.ID
	result.SetSuccess()
	return result, nil
}

// UpdateStruct 更新结构体
func UpdateStruct(reqInfo model.UpdateStructReq) (model.UpdateStructResp, error) {
	var result model.UpdateStructResp

	tx := GetDB().Begin()
	data := reqInfo.DBStructInfo

	if err := tx.Where("id = ?", data.ID).Save(data).Error; err != nil {
		tx.Rollback()
		log.Error("更新 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, f := range data.Fields {
		f.ParentID = data.ID
		if err := tx.Save(&f).Error; err != nil {
			tx.Rollback()
			log.Error("保存 结构体的字段 信息错误:%s", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	tx.Commit()
	result.ID = data.ID
	result.SetSuccess()
	return result, nil
}

// DeleteStruct 删除结构体
func DeleteStruct(id int) (model.DeleteStructResp, error) {
	var result model.DeleteStructResp

	tx := GetDB().Begin()

	if err := tx.Where("id = ?", id).Delete(&model.DBStructInfo{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := tx.Where("parent_id = ?", id).Delete(&model.DBField{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 结构体对应字段 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	tx.Commit()
	result.ID = id
	result.SetSuccess()
	return result, nil
}
