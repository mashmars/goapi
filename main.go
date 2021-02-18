package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"api/route"
	"fmt"
)


func ErrHandler() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(200, gin.H{
					"code": 1,
					"msg" : fmt.Sprintf("%v", err),
				})
				return
			}
		}()
		//这个必须有
		ctx.Next()
	}
}

func main() {
	
	router := gin.Default()
	//解决跨域请求
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}//允许请求的方法
	//allowheaders需要设置上 Access-Control-Allow-Origin 
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*", "Access-Control-Allow-Origin"}//允许的Header
	router.Use(cors.New(config))
	//router.Use(cors.Default())
	router.Use(ErrHandler())

	//注册路由
	route.LoadSecurity(router)
	route.LoadAdmin(router)
	route.LoadAdminRole(router)

	router.Run()
}