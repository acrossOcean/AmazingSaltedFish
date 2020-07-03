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
// @Tags project定义信息,获取,定义
// @Accept json
// @Produce json
// @Param id path int true "ID" minimum(1)
// @Success 200 {object} model.GetProjectDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /project/define/detail/{id} [get]
func GetProjectDefine(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.GetProjectDefine(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 获取结构信息列表
// @Tags project定义信息,获取,列表,定义
// @Accept json
// @Produce json
// @Param pageNum query int true "分页信息, 页数从1开始"
// @Param pageSize query int true "分页信息, 每页最大500"
// @Success 200 {object} model.GetProjectDefineListResp
// @Failure 400 {object} model.BaseResp
// @Router /project/define/list [get]
func GetProjectDefineList(ctx *gin.Context) {
	reqInfo := getPageInfo(ctx)

	if reqInfo.PageNum < 1 {
		reqInfo.PageNum = 1
	}

	if reqInfo.PageSize < 0 || reqInfo.PageSize > 500 {
		reqInfo.PageSize = 20
	}

	info, err := service.GetProjectDefineList(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 新建一个结构信息
// @Tags project定义信息,新建,定义
// @Accept json
// @Produce json
// @Param projectInfo body model.CreateProjectDefineReq true "结构体信息"
// @Success 200 {object} model.CreateProjectDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /project/define/ [post]
func CreateProjectDefine(ctx *gin.Context) {
	var reqInfo model.CreateProjectDefineReq
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

	info, err := service.CreateProjectDefine(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 更新一个结构信息
// @Tags project定义信息,更新,定义
// @Accept json
// @Produce json
// @Param projectInfo body model.UpdateProjectDefineReq true "结构体信息"
// @Success 200 {object} model.UpdateProjectDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /project/define/ [put]
func UpdateProjectDefine(ctx *gin.Context) {
	var reqInfo model.UpdateProjectDefineReq
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

	info, err := service.UpdateProjectDefine(reqInfo)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}

// @Summary 删除一个结构信息
// @Tags project定义信息,删除,定义
// @Accept json
// @Produce json
// @Param id path int true "project的ID" minimum(1)
// @Success 200 {object} model.DeleteProjectDefineResp
// @Failure 400 {object} model.BaseResp
// @Router /project/define/{id} [delete]
func DeleteProjectDefine(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if !check.PassCheck(
		check.NewIntChecker(id, check.NewIntCheckOption(check.IntOperatorGT, 0)),
	) {
		respParamError(ctx, id)
		return
	}

	info, err := service.DeleteProjectDefine(id)
	if err != nil {
		respError(ctx, err)
		return
	}

	respSuccessWithInfo(ctx, info)
}
