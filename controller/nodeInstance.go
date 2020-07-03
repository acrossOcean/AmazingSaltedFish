package controller

import (
	"AmazingSaltedFish/model"
	"AmazingSaltedFish/service"
	"AmazingSaltedFish/utils/check"
	"strconv"

	"github.com/acrossOcean/log"

	"github.com/gin-gonic/gin"
)

// GetNodeInstance 操作 node instance api
// @Summary 获取一个结构信息
// @Tags node实体信息,获取,实体
// @Accept json
// @Produce json
// @Param id path int true "ID" minimum(1)
// @Success 200 {object} model.GetNodeInstanceResp
// @Failure 400 {object} model.BaseResp
// @Router /node/instance/detail/{id} [get]
func GetNodeInstance(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.GetNodeInstance(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// GetNodeInstanceList 操作 node instance api
// @Summary 获取结构信息列表
// @Tags node实体信息,获取,列表,实体
// @Accept json
// @Produce json
// @Param pageNum query int true "分页信息, 页数从1开始"
// @Param pageSize query int true "分页信息, 每页最大500"
// @Success 200 {object} model.GetNodeInstanceListResp
// @Failure 400 {object} model.BaseResp
// @Router /node/instance/list [get]
func GetNodeInstanceList(ctx *gin.Context) {
	reqInfo := getPageInfo(ctx)

	if reqInfo.PageNum < 1 {
		reqInfo.PageNum = 1
	}

	if reqInfo.PageSize < 0 || reqInfo.PageSize > 500 {
		reqInfo.PageSize = 20
	}

	info, err := service.GetNodeInstanceList(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// CreateNodeInstance 操作 node instance api
// @Summary 新建一个结构信息
// @Tags node实体信息,新建,实体
// @Accept json
// @Produce json
// @Param projectInfo body model.CreateNodeInstanceReq true "结构体信息"
// @Success 200 {object} model.CreateNodeInstanceResp
// @Failure 400 {object} model.BaseResp
// @Router /node/instance/ [post]
func CreateNodeInstance(ctx *gin.Context) {
	var reqInfo model.CreateNodeInstanceReq
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

	for _, p := range reqInfo.ParamList {
		if !check.PassCheck(
			check.NewIntChecker(p.PType.ToInt(),
				check.NewIntCheckOptionIN(model.GetAllFieldTypeInt()),
			),
			check.NewStrChecker(p.Name, check.NewStrCheckOption(check.StrOperatorLenLE, 20, "")),
			check.NewStrChecker(p.Comment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
		) {
			respParamError(ctx, reqInfo)
			return
		}
	}

	info, err := service.CreateNodeInstance(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// UpdateNodeInstance 操作 node instance api
// @Summary 更新一个结构信息
// @Tags node实体信息,更新,实体
// @Accept json
// @Produce json
// @Param projectInfo body model.UpdateNodeInstanceReq true "结构体信息"
// @Success 200 {object} model.UpdateNodeInstanceResp
// @Failure 400 {object} model.BaseResp
// @Router /node/instance/ [put]
func UpdateNodeInstance(ctx *gin.Context) {
	var reqInfo model.UpdateNodeInstanceReq
	if err := ctx.ShouldBind(&reqInfo); err != nil {
		log.Debug("接收参数错误:%s", err.Error())
		respParamError(ctx, reqInfo)
		return
	}

	if !check.PassCheck(
		check.NewIntChecker(reqInfo.ID, check.NewIntCheckOption(check.IntOperatorGT, 0)),
		check.NewStrChecker(reqInfo.Comment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
	) {
		respParamError(ctx, reqInfo)
		return
	}

	for _, p := range reqInfo.ParamList {
		if !check.PassCheck(
			check.NewIntChecker(p.PType.ToInt(),
				check.NewIntCheckOptionIN(model.GetAllFieldTypeInt()),
			),
			check.NewStrChecker(p.Name, check.NewStrCheckOption(check.StrOperatorLenLE, 20, "")),
			check.NewStrChecker(p.Comment, check.NewStrCheckOption(check.StrOperatorLenLE, 200, "")),
		) {
			respParamError(ctx, reqInfo)
			return
		}
	}

	info, err := service.UpdateNodeInstance(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// DeleteNodeInstance 操作 node instance api
// @Summary 删除一个结构信息
// @Tags node实体信息,删除,实体
// @Accept json
// @Produce json
// @Param id path int true "project的ID" minimum(1)
// @Success 200 {object} model.DeleteNodeInstanceResp
// @Failure 400 {object} model.BaseResp
// @Router /node/instance/{id} [delete]
func DeleteNodeInstance(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.DeleteNodeInstance(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}
