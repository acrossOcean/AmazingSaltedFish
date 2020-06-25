package controller

import "github.com/gin-gonic/gin"

// @Summary 获取一个结构信息
// @Tags 结构体信息,获取
// @Accept json
// @Produce json
// @Param structId path int true "struct的ID" minimum(1)
// @Success 200 {object} model.GetStructResp
// @Failure 400 {object} model.BaseResp
// @Router /detail/{id} [get]
func GetStruct(ctx *gin.Context) {
}

// @Summary 新建一个结构信息
// @Tags 结构体信息,新建
// @Accept json
// @Produce json
// @Param structInfo body model.CreateStructReq true "结构体信息"
// @Success 200 {object} model.CreateStructResp
// @Failure 400 {object} model.BaseResp
// @Router / [post]
func CreateStruct(ctx *gin.Context) {
}

// @Summary 更新一个结构信息
// @Tags 结构体信息,更新
// @Accept json
// @Produce json
// @Param structInfo body model.UpdateStructReq true "结构体信息"
// @Success 200 {object} model.UpdateStructResp
// @Failure 400 {object} model.BaseResp
// @Router / [put]
func UpdateStruct(ctx *gin.Context) {
}

// @Summary 删除一个结构信息
// @Tags 结构体信息,删除
// @Accept json
// @Produce json
// @Param structId path int true "struct的ID" minimum(1)
// @Success 200 {object} model.DeleteStructResp
// @Failure 400 {object} model.BaseResp
// @Router /{id} [delete]
func DeleteStruct(ctx *gin.Context) {
}

// @Summary 获取结构信息列表
// @Tags 结构体信息,获取,列表
// @Accept json
// @Produce json
// @Param pageNum query int true "页数" minimum(1)
// @Param pageSize query int true "每页大小" Enums(10,20,50,100)
// @Success 200 {object} model.GetStructListResp
// @Failure 400 {object} model.BaseResp
// @Router /list [get]
func GetStructList(ctx *gin.Context) {
}
