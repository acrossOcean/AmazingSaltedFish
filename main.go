package main

import (
	"os"

	"AmazingSaltedFish/constant"
	"AmazingSaltedFish/docs"
	"AmazingSaltedFish/router"
	"AmazingSaltedFish/service"

	"github.com/acrossOcean/config"
	"github.com/acrossOcean/log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	log.SetStaticTags(map[string]interface{}{
		"_appName":    constant.ConstAppName,
		"_appVersion": constant.ConstAppVersion,
	})
}

// @title Amazing Salted Fish API
// @version 0.0.1
// @description this is amazing salted fish swagger api page.
// @termsOfService http://www.amazingsaltedfish.com

// @contact.name guoyueyang
// @contact.url http://www.amazingsaltedfish.com
// @contact.email guoyueyang@126.com

// @license.name MIT
// @license.url https://spdx.org/licenses/MIT.html
func main() {
	// 初始化配置文件
	config.AddPath("./config.ini")
	log.Debug(config.Get("mysql"))
	log.Debug(config.GetCurrentCache())

	// 初始化数据库
	err := service.InitDB()
	if err != nil {
		log.Error("初始化数据库错误:%s", err.Error())
		os.Exit(1)
	}
	defer log.Error(service.CloseDB())

	engine := gin.Default()
	r := engine.Group("/api")
	router.InitRouter(r)

	// swagger 支持
	if config.DefaultBool("swagger>>enableSwagger", false) {
		docs.SwaggerInfo.Host = config.DefaultString("main>>listenAddr", ":8080")
		docs.SwaggerInfo.BasePath = "/api/v1"
		engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, config.DefaultString("swagger>>closeEnvName", "ENABLE_SWAGGER")))
	}

	log.Error(engine.Run(config.DefaultString("main>>listenAddr", ":8080")))
}