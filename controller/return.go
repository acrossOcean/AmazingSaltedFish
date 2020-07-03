package controller

import (
	"AmazingSaltedFish/utils"
	"net/http"

	"AmazingSaltedFish/model"

	"github.com/acrossOcean/log"
	"github.com/gin-gonic/gin"
)

func respParamError(ctx *gin.Context, param interface{}) {
	log.Debug("请求发送参数错误, 参数为:%v", utils.GetJSONOutput(param))
	ctx.JSON(http.StatusBadRequest, model.RespParamError)
}

func respError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, model.NewBaseRespWithError(err))
}

//func respSuccess(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, model.RespSuccess)
//}

func respSuccessWithInfo(ctx *gin.Context, info interface{}) {
	ctx.JSON(http.StatusOK, info)
}
