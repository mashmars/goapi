package securitycontroller

import (
	"github.com/gin-gonic/gin"
	"api/model"
)

func Login(ctx *gin.Context) {
	var adminModel model.Admin
	//find 根据主键返回row
	admin := adminModel.Find(1)
	
	ctx.JSON(200, gin.H{
		"code": 0, 
		"msg" : "登陆成功",
		"data": admin,
	})
}