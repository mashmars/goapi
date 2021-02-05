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
	router.Use(cors.Default())
	router.Use(ErrHandler())

	//注册路由
	route.LoadSecurity(router)
	route.LoadAdmin(router)

	router.Run()
}