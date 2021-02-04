package securitycontroller

import (
	"github.com/gin-gonic/gin"
	//"api/model"
)

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 0, 
		"msg" : "登陆成功",
		"data": "",
	})
}