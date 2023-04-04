package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunanxiang/charlotter/cache"
	"github.com/sunanxiang/charlotter/server"
)

func main() {
	router := gin.New()
	router.POST("/wechat", server.HandleWechat)
	router.GET("/wechat", server.Verify)

	cache.Init()
	// 运行服务器
	router.Run(":80")
}
