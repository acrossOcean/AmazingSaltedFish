package controller

import (
	"AmazingSaltedFish/controller/check"
	"AmazingSaltedFish/model"
	"AmazingSaltedFish/service"
	"strconv"

	"github.com/acrossOcean/log"

	"github.com/gin-gonic/gin"
)

// @Summary 获取一个结构信息
// @Tags 结构体信息,获取
// @Accept json
// @Produce json
// @Param id path int true "struct的ID" minimum(1)
// @Success 200 {object} model.GetStructResp
// @Failure 400 {object} model.BaseResp
// @Router /struct/detail/{id} [get]
func GetStruct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.GetStruct(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 新建一个结构信息
// @Tags 结构体信息,新建
// @Accept json
// @Produce json
// @Param structInfo body model.CreateStructReq true "结构体信息"
// @Success 200 {object} model.CreateStructResp
// @Failure 400 {object} model.BaseResp
// @Router /struct/ [post]
func CreateStruct(ctx *gin.Context) {
	var reqInfo model.CreateStructReq
	if err := ctx.ShouldBind(&reqInfo); err != nil {
		log.Debug("接收参数错误:%s", err.Error())
		respParamError(ctx, reqInfo)
		return
	}

	if !check.PassCheck(
		check.NewStrChecker(reqInfo.Comment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
	) {
		respParamError(ctx, reqInfo)
		return
	}

	for _, f := range reqInfo.Fields {
		if !check.PassCheck(
			check.NewIntChecker(f.FType.ToInt(),
				check.NewIntCheckOptionIN([]int{
					model.FieldTypeString.ToInt(),
					model.FieldTypeBool.ToInt(),
					model.FieldTypeInt.ToInt(),
					model.FieldTypeFloat.ToInt(),
					model.FieldTypeStruct.ToInt()}),
			),
			check.NewStrChecker(f.FName, check.NewStrCheckOption(check.StrOperatorLenLE, 20, "")),
			check.NewStrChecker(f.FComment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
		) {
			respParamError(ctx, reqInfo)
			return
		}
	}

	info, err := service.CreateStruct(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 更新一个结构信息
// @Tags 结构体信息,更新
// @Accept json
// @Produce json
// @Param structInfo body model.UpdateStructReq true "结构体信息"
// @Success 200 {object} model.UpdateStructResp
// @Failure 400 {object} model.BaseResp
// @Router /struct/ [put]
func UpdateStruct(ctx *gin.Context) {
	var reqInfo model.UpdateStructReq
	if err := ctx.ShouldBind(&reqInfo); err != nil {
		log.Debug("接收参数错误:%s", err.Error())
		respParamError(ctx, reqInfo)
		return
	}

	if !check.PassCheck(
		check.NewIntChecker(reqInfo.Id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
		check.NewStrChecker(reqInfo.Comment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
	) {
		respParamError(ctx, reqInfo)
		return
	}

	for _, f := range reqInfo.Fields {
		if !check.PassCheck(
			check.NewIntChecker(f.FType.ToInt(),
				check.NewIntCheckOptionIN([]int{
					model.FieldTypeString.ToInt(),
					model.FieldTypeBool.ToInt(),
					model.FieldTypeInt.ToInt(),
					model.FieldTypeFloat.ToInt(),
					model.FieldTypeStruct.ToInt()}),
			),
			check.NewStrChecker(f.FName, check.NewStrCheckOption(check.StrOperatorLenLE, 20, "")),
			check.NewStrChecker(f.FComment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
		) {
			respParamError(ctx, reqInfo)
			return
		}
	}

	info, err := service.UpdateStruct(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 删除一个结构信息
// @Tags 结构体信息,删除
// @Accept json
// @Produce json
// @Param id path int true "struct的ID" minimum(1)
// @Success 200 {object} model.DeleteStructResp
// @Failure 400 {object} model.BaseResp
// @Router /struct/{id} [delete]
func DeleteStruct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.DeleteStruct(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 获取结构信息列表
// @Tags 结构体信息,获取,列表
// @Accept json
// @Produce json
// @Param pageNum query int true "分页信息, 页数从1开始"
// @Param pageSize query int true "分页信息, 每页最大500"
// @Success 200 {object} model.GetStructListResp
// @Failure 400 {object} model.BaseResp
// @Router /struct/list [get]
func GetStructList(ctx *gin.Context) {
	reqInfo := getPageInfo(ctx)

	if reqInfo.PageNum < 1 {
		reqInfo.PageNum = 1
	}

	if reqInfo.PageSize < 0 || reqInfo.PageSize > 500 {
		reqInfo.PageSize = 20
	}

	info, err := service.GetStructList(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}
