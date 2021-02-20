package adminrolecontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
	"errors"
	"gorm.io/gorm"
)

func All(ctx *gin.Context) {
	var adminRole model.AdminRole1
	adminRoles := adminRole.FindAll()
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminRoles,
	})
}


func Index(ctx *gin.Context) {
	var adminRoles []model.AdminRole
	if result := model.ORM.Find(&adminRoles); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminRoles,
	})
}

func Add(ctx *gin.Context) {
	var adminRole model.AdminRole
	if err := ctx.ShouldBind(&adminRole); err != nil {
		panic(err)
	}
	if result := model.ORM.Create(&adminRole); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Edit(ctx *gin.Context) {
	id := ctx.Param("id")
	var adminRole model.AdminRole
	model.ORM.First(&adminRole, id)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminRole,
	})
}

func EditSave(ctx *gin.Context) {
	id := ctx.Param("id")
	var adminRole model.AdminRole
	model.ORM.First(&adminRole, id)

	if err := ctx.ShouldBind(&adminRole); err != nil {
		panic(err)
	}

	if result := model.ORM.Updates(&adminRole); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Status(ctx *gin.Context) {
	id := ctx.DefaultPostForm("id", "0")
	status := ctx.PostForm("status")
	var adminRole model.AdminRole
	if result := model.ORM.First(&adminRole, id); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		
	}
	if result := model.ORM.Model(&adminRole).Update("is_enabled", status); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.DefaultPostForm("id", "0")
	var adminRole model.AdminRole
	if result := model.ORM.First(&adminRole, id); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		
	}
	model.ORM.Delete(&adminRole)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}