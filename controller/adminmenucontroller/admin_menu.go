package adminmenucontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	var adminMenus []model.AdminMenu
	model.ORM.Find(&adminMenus)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminMenus,
	})
}

func Add(ctx *gin.Context) {
	var adminMenu model.AdminMenu
	if err := ctx.ShouldBind(&adminMenu); err != nil {
		panic(err)
	}
	if result := model.ORM.Create(&adminMenu); result.Error != nil {
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
	var adminMenu model.AdminMenu
	model.ORM.Find(&adminMenu, id)
	/*if result := model.ORM.First(&adminMenu, id); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg" : "请求信息不存在",
			"data": "",
		})
		return
	}*/
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminMenu,
	})
}

func EditSave(ctx *gin.Context) {
	id := ctx.Param("id")
	var adminMenu model.AdminMenu
	model.ORM.Find(&adminMenu, id)

	if err := ctx.ShouldBind(&adminMenu); err != nil {
		panic(err)
	}

	if result := model.ORM.Save(&adminMenu); result.Error != nil {
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
	var adminMenu model.AdminMenu
	model.ORM.Find(&adminMenu, id)
	
	if result := model.ORM.Model(&adminMenu).Update("is_enabled", status); result.Error != nil {
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
	var adminMenu model.AdminMenu
	model.ORM.Find(&adminMenu, id)
	
	if result := model.ORM.Delete(&adminMenu); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func MenuAction(ctx *gin.Context) {
	id := ctx.Param("id")
	var adminActions []model.AdminAction
	model.ORM.Where("menu_id = ?", id).Find(&adminActions)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminActions,
	})
}

func All(ctx *gin.Context) {
	var adminMenus []model.AdminMenu
	model.ORM.Find(&adminMenus)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminMenus,
	})
}