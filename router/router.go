package router

import (
	"AmazingSaltedFish/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	v1Router := r.Group("/v1")

	// 添加框架
	InitStructRouter(v1Router)
}

func InitStructRouter(router *gin.RouterGroup) {
	r := router.Group("/struct")
	{
		r.POST("/", controller.CreateStruct)
		r.PUT("/", controller.UpdateStruct)
		r.GET("/:id", controller.GetStruct)
		r.DELETE("/:id", controller.DeleteStruct)

		r.GET("/list", controller.GetStructList)
	}
}
