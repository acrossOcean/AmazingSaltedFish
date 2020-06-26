package service

import (
	"AmazingSaltedFish/model"
	"sort"

	"github.com/jinzhu/gorm"

	"github.com/acrossOcean/log"
)

// 获取结构体详情
func GetStruct(id int) (model.GetStructResp, error) {
	var result model.GetStructResp
	var info model.StructInfo

	err := _DB.Where("id = ?", id).First(&info).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeResourceNotExist, err)
	} else if err != nil {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	fields := make([]model.Field, 0)
	err = _DB.Where("parent_id = ?", info.Id).Find(&fields).Error
	if err != nil {
		log.Error("根据 结构体ID(%v) 获取 对应字段信息 错误:%s", info.Id, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}
	sort.Sort(model.FieldList(fields))

	info.Fields = fields
	result.Info = info
	result.SetSuccess()

	return result, nil
}

// 新建结构体
func CreateStruct(reqInfo model.CreateStructReq) (model.CreateStructResp, error) {
	var result model.CreateStructResp

	tx := _DB.Begin()
	data := reqInfo.ToNormal()

	if err := tx.Save(&data).Error; err != nil {
		tx.Rollback()
		log.Error("保存 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, f := range data.Fields {
		f.ParentId = data.Id
		if err := tx.Save(&f).Error; err != nil {
			tx.Rollback()
			log.Error("保存 结构体的字段 信息错误:%s", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	tx.Commit()
	result.Id = data.Id
	return result, nil
}

// 更新结构体
func UpdateStruct(reqInfo model.UpdateStructReq) (model.UpdateStructResp, error) {
	var result model.UpdateStructResp

	tx := _DB.Begin()
	data := reqInfo.StructInfo

	if err := tx.Where("id = ?", data.Id).Save(data).Error; err != nil {
		tx.Rollback()
		log.Error("更新 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	// 删除之前的字段,再新建新的字段
	if err := tx.Where("parent_id = ?", data.Id).Delete(&model.Field{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 结构体对应字段 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	for _, f := range data.Fields {
		f.ParentId = data.Id
		if err := tx.Save(&f).Error; err != nil {
			tx.Rollback()
			log.Error("保存 结构体的字段 信息错误:%s", err.Error())
			return result, model.ReturnWithCode(model.CodeUnknownError, err)
		}
	}

	tx.Commit()
	result.Id = data.Id
	return result, nil
}

// 删除结构体
func DeleteStruct(id int) (model.DeleteStructResp, error) {
	var result model.DeleteStructResp

	tx := _DB.Begin()

	if err := tx.Where("id = ?", id).Delete(&model.StructInfo{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := tx.Where("parent_id = ?", id).Delete(&model.Field{}).Error; err != nil {
		tx.Rollback()
		log.Error("删除 结构体对应字段 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	tx.Commit()
	result.Id = id
	return result, nil
}

// 获取结构体列表
func GetStructList(req model.PageReq) (model.GetStructListResp, error) {
	var result model.GetStructListResp
	var infos []model.StructInfo
	var count int

	if err := _DB.Table(model.StructInfo{}.TableName()).Count(&count).Error; err != nil {
		log.Error("获取 结构体表 数据总数错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	if err := _DB.Order("id DESC").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&infos).Error; err != nil {
		log.Error("根据 id 获取 结构体 信息错误:%s", err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	structIds := make([]int, len(infos))
	for i, info := range infos {
		structIds[i] = info.Id
	}

	fields := make([]model.Field, 0)
	if err := _DB.Where("parent_id IN (?)", structIds).Find(&fields).Error; err != nil {
		log.Error("根据 结构体IDs(%v) 获取 对应字段信息 错误:%s", structIds, err.Error())
		return result, model.ReturnWithCode(model.CodeUnknownError, err)
	}

	fieldMap := make(map[int][]model.Field)
	for _, f := range fields {
		if _, ok := fieldMap[f.ParentId]; !ok {
			fieldMap[f.ParentId] = make([]model.Field, 0)
		}

		fieldMap[f.ParentId] = append(fieldMap[f.ParentId], f)
	}

	for i, info := range infos {
		if list, ok := fieldMap[info.Id]; ok {
			sort.Sort(model.FieldList(list))
			info.Fields = list
			infos[i] = info
		} else {
			info.Fields = make([]model.Field, 0)
			infos[i] = info
		}
	}

	result.SUm = count
	result.List = infos
	result.SetSuccess()

	return result, nil
}
