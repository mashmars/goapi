package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"api/route"
)

func main() {
	router := gin.Default()
	//解决跨域请求
	router.Use(cors.Default())

	//注册路由
	route.LoadSecurity(router)

	router.Run()
}