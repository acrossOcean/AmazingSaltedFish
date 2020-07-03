package router

import (
	"AmazingSaltedFish/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(r *gin.RouterGroup) {
	v1Router := r.Group("/v1")

	// 添加框架
	initStructRouter(v1Router)
	initProjectRouter(v1Router)
	initLinkRouter(v1Router)
	initNodeRouter(v1Router)
}

// 结构定义
func initStructRouter(router *gin.RouterGroup) {
	r := router.Group("/struct")
	{
		r.POST("/", controller.CreateStruct)
		r.PUT("/", controller.UpdateStruct)
		r.DELETE("/:id", controller.DeleteStruct)

		r.GET("/detail/:id", controller.GetStruct)
		r.GET("/list", controller.GetStructList)
	}
}

// 项目路由
func initProjectRouter(router *gin.RouterGroup) {
	r := router.Group("/proj")
	{
		defineRouter := r.Group("/define")
		{
			// 获取项目/列表
			defineRouter.GET("/detail/:id", controller.GetProjectDefine)
			defineRouter.GET("/list", controller.GetProjectDefineList)
			// 新建项目定义
			defineRouter.POST("/", controller.CreateProjectDefine)
			// 更新项目定义
			defineRouter.PUT("/", controller.UpdateProjectDefine)
			// 删除项目定义 (会同步删除实现信息)
			defineRouter.DELETE("/", controller.DeleteProjectDefine)
		}

		instanceRouter := r.Group("/instance")
		{
			// 获取项目/列表
			instanceRouter.GET("/detail/:id", controller.GetProjectInstance)
			instanceRouter.GET("/list", controller.GetProjectInstanceList)
			// 新建项目实现
			instanceRouter.POST("/", controller.CreateProjectInstance)
			// 更新项目实现
			instanceRouter.PUT("/", controller.UpdateProjectInstance)
			// 删除项目实现 (保留项目定义)
			instanceRouter.DELETE("/", controller.DeleteProjectInstance)
		}
	}
}

// link 路由
func initLinkRouter(router *gin.RouterGroup) {
	r := router.Group("/link")
	{
		defineRouter := r.Group("/define")
		{
			// 获取项目/列表
			defineRouter.GET("/detail/:id", controller.GetLinkDefine)
			defineRouter.GET("/list", controller.GetLinkDefineList)
			// 新建项目定义
			defineRouter.POST("/", controller.CreateLinkDefine)
			// 更新项目定义
			defineRouter.PUT("/", controller.UpdateLinkDefine)
			// 删除项目定义 (会同步删除实现信息)
			defineRouter.DELETE("/", controller.DeleteLinkDefine)
		}

		instanceRouter := r.Group("/instance")
		{
			// 获取项目/列表
			instanceRouter.GET("/detail/:id", controller.GetLinkInstance)
			instanceRouter.GET("/list", controller.GetLinkInstanceList)
			// 新建项目实现
			instanceRouter.POST("/", controller.CreateLinkInstance)
			// 更新项目实现
			instanceRouter.PUT("/", controller.UpdateLinkInstance)
			// 删除项目实现 (保留项目定义)
			instanceRouter.DELETE("/", controller.DeleteLinkInstance)
		}
	}
}

// node 路由
func initNodeRouter(router *gin.RouterGroup) {
	r := router.Group("/node")
	{
		defineRouter := r.Group("/define")
		{
			// 获取项目/列表
			defineRouter.GET("/detail/:id", controller.GetNodeDefine)
			defineRouter.GET("/list", controller.GetNodeDefineList)
			// 新建项目定义
			defineRouter.POST("/", controller.CreateNodeDefine)
			// 更新项目定义
			defineRouter.PUT("/", controller.UpdateNodeDefine)
			// 删除项目定义 (会同步删除实现信息)
			defineRouter.DELETE("/", controller.DeleteNodeDefine)
		}

		instanceRouter := r.Group("/instance")
		{
			// 获取项目/列表
			instanceRouter.GET("/detail/:id", controller.GetNodeInstance)
			instanceRouter.GET("/list", controller.GetNodeInstanceList)
			// 新建项目实现
			instanceRouter.POST("/", controller.CreateNodeInstance)
			// 更新项目实现
			instanceRouter.PUT("/", controller.UpdateNodeInstance)
			// 删除项目实现 (保留项目定义)
			instanceRouter.DELETE("/", controller.DeleteNodeInstance)
		}
	}
}
