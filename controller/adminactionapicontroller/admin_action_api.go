package adminactionapicontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
)

func Api(ctx *gin.Context) {
	var adminActionApis []model.AdminActionApi
	model.ORM.Where("is_enabled = 1").Order("controller_action asc").Find(&adminActionApis)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminActionApis,
	})
}


func ApiSetAction(ctx *gin.Context) {
	id := ctx.PostForm("action_id") //action id
	ids := ctx.PostFormArray("ids[]") //api id
	
	model.ORM.Model(&model.AdminActionApi{}).Where("action_id = ?", id).Update("action_id", 0)
	model.ORM.Model(&model.AdminActionApi{}).Where("id in (?)", ids).Update("action_id", id) 
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}