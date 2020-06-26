package controller

import (
	"AmazingSaltedFish/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPageInfo(ctx *gin.Context) model.PageReq {
	var result model.PageReq

	sizeStr := ctx.Query("pageSize")
	size, _ := strconv.Atoi(sizeStr)

	numStr := ctx.Query("pageNum")
	num, _ := strconv.Atoi(numStr)

	result.PageSize = size
	result.PageNum = num

	return result
}
