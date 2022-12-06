package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/* 常用命令
go get :下载依赖
go get -u:升级依赖到最新
go get -u=patch:升级依赖到最新的修订版
go get 地址@版本号：升级依赖到某个版本
go mod tidy: 整理依赖
go mod init：初始化mod文件
*/
