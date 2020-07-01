package main

import (
	"os"
	"strings"

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

	log.SetLevel(log.ConvertToLevel(config.DefaultInt("log>>level", 0)))
}

// @title Amazing Salted Fish API
// @version 0.0.1
// @description this is [amazing salted fish] swagger api page.
// @termsOfService http://www.amazingsaltedfish.com

// @contact.name author:guoyueyang
// @contact.url http://www.amazingsaltedfish.com
// @contact.email guoyueyang@126.com

// @license.name License:MIT
// @license.url https://spdx.org/licenses/MIT.html
func main() {
	// 初始化配置文件
	config.AddPath("./config.ini")

	// 初始化数据库
	err := service.InitDB()
	if err != nil {
		log.Error("初始化数据库错误:%s", err.Error())
		os.Exit(1)
	}
	defer service.CloseDB()

	engine := gin.Default()
	r := engine.Group("/api")
	router.InitRouter(r)

	// swagger 支持
	log.Debug("是否开启swagger:%v", config.DefaultBool("swagger>>enableSwagger", false))
	if config.DefaultBool("swagger>>enableSwagger", false) {
		addr := config.DefaultString("main>>listenAddr", ":8080")
		if strs := strings.Split(addr, ":"); len(strs) == 2 {
			if strs[0] == "0.0.0.0" || strs[0] == "" {
				addr = "localhost" + ":" + strs[1]
			}
		}
		docs.SwaggerInfo.Host = addr
		docs.SwaggerInfo.BasePath = "/api/v1"
		engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, config.DefaultString("swagger>>closeEnvName", "ENABLE_SWAGGER")))
	}

	log.Error(engine.Run(config.DefaultString("main>>listenAddr", ":8080")))
}
