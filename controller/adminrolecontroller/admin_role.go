package adminrolecontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context) {
	var adminRole model.AdminRole
	adminRoles := adminRole.FindAll()
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminRoles,
	})
}