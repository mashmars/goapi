package adminrolecontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"math"
)

func All(ctx *gin.Context) {
	var adminRoles []model.AdminRole
	model.ORM.Find(&adminRoles)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminRoles,
	})
}


func Index(ctx *gin.Context) {
	var adminRoles []model.AdminRole
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"));
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	type QueryForm struct {
		Name string `form:"name"`
		IsEnabled string `form:"is_enabled"`
	}

	var queryForm QueryForm
	if ctx.ShouldBindQuery(&queryForm) != nil {
		panic("请求参数绑定失败")
	}	

	offset := (page - 1) * pageSize
	var count int64
	model.ORM.Model(&model.AdminRole{}).Where(queryForm).Count(&count)
	pages := math.Ceil(float64(count)/float64(pageSize))

	if result := model.ORM.Where(queryForm).Limit(pageSize).Offset(offset).Find(&adminRoles); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminRoles,
		"totalPage": pages,
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

	if result := model.ORM.Model(&adminRole).Select("*").Omit("ID").Updates(adminRole); result.Error != nil {
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