package controller

import (
	"AmazingSaltedFish/model"
	"AmazingSaltedFish/service"
	"AmazingSaltedFish/utils/check"
	"strconv"

	"github.com/acrossOcean/log"

	"github.com/gin-gonic/gin"
)

// @Summary 获取一个结构信息
// @Tags node定义信息,获取,定义
// @Accept json
// @Produce json
// @Param id path int true "ID" minimum(1)
// @Success 200 {object} model.GetNodeDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /node/define/detail/{id} [get]
func GetNodeDefine(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIDChecker(id),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.GetNodeDefine(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 获取结构信息列表
// @Tags node定义信息,获取,列表,定义
// @Accept json
// @Produce json
// @Param pageNum query int true "分页信息, 页数从1开始"
// @Param pageSize query int true "分页信息, 每页最大500"
// @Success 200 {object} model.GetNodeDefineListResp
// @Failure 400 {object} model.BaseResp
// @Router /node/define/list [get]
func GetNodeDefineList(ctx *gin.Context) {
	reqInfo := getPageInfo(ctx)

	info, err := service.GetNodeDefineList(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 新建一个结构信息
// @Tags node定义信息,新建,定义
// @Accept json
// @Produce json
// @Param projectInfo body model.CreateNodeDefineReq true "结构体信息"
// @Success 200 {object} model.CreateNodeDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /node/define/ [post]
func CreateNodeDefine(ctx *gin.Context) {
	var reqInfo model.CreateNodeDefineReq
	if err := ctx.ShouldBind(&reqInfo); err != nil {
		log.Debug("接收参数错误:%s", err.Error())
		respParamError(ctx, reqInfo)
		return
	}

	if !reqInfo.Check() {
		respParamError(ctx, reqInfo)
		return
	}

	info, err := service.CreateNodeDefine(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 更新一个结构信息
// @Tags node定义信息,更新,定义
// @Accept json
// @Produce json
// @Param projectInfo body model.UpdateNodeDefineReq true "结构体信息"
// @Success 200 {object} model.UpdateNodeDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /node/define/ [put]
func UpdateNodeDefine(ctx *gin.Context) {
	var reqInfo model.UpdateNodeDefineReq
	if err := ctx.ShouldBind(&reqInfo); err != nil {
		log.Debug("接收参数错误:%s", err.Error())
		respParamError(ctx, reqInfo)
		return
	}

	if !reqInfo.Check() {
		respParamError(ctx, reqInfo)
		return
	}

	info, err := service.UpdateNodeDefine(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 删除一个结构信息
// @Tags node定义信息,删除,定义
// @Accept json
// @Produce json
// @Param id path int true "project的ID" minimum(1)
// @Success 200 {object} model.DeleteNodeDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /node/define/{id} [delete]
func DeleteNodeDefine(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.DeleteNodeDefine(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}
