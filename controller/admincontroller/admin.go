package admincontroller

import (
	"github.com/gin-gonic/gin"
	"api/model"
	"strconv"
	"api/service/passwordencoder/argon2idencoder"
	"math"
)

func Index(ctx *gin.Context) {
	var admins []model.Admin
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"));
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))	

	offset := (page - 1) * pageSize
	var count int64
	model.ORM.Model(&model.Admin{}).Count(&count)
	pages := math.Ceil(float64(count)/float64(pageSize))

	model.ORM.Select("admin.*, admin_role.name as role_name").Joins("left join admin_role on admin_role.id = admin.role_id").Limit(pageSize).Offset(offset).Find(&admins)
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": admins,
		"totalPage": pages,
	})
}

func Add(ctx *gin.Context) {	
	var admin model.Admin
	if err := ctx.ShouldBind(&admin); err != nil {
		panic(err)
	}	

	if result := model.ORM.Create(&admin); result.Error != nil {
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
	var admin model.Admin
	model.ORM.Find(&admin, id)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": admin,
	})
}

func EditSave(ctx *gin.Context) {
	id := ctx.Param("id")
	var admin model.Admin
	model.ORM.Find(&admin, id)

	if err := ctx.ShouldBind(&admin); err != nil {
		panic(err)
	}
	if result := model.ORM.Save(&admin); result.Error != nil {
		panic(result.Error)
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var admin model.Admin
	model.ORM.Find(&admin, id)
	model.ORM.Delete(&admin)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Status(ctx *gin.Context) {
	id := ctx.DefaultPostForm("id", "0")
	status := ctx.PostForm("status")
	var admin model.Admin
	model.ORM.Find(&admin, id)
	model.ORM.Model(&admin).Update("is_enabled", status)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Password(ctx *gin.Context) {
	old_password := ctx.PostForm("old_password")
	password := ctx.PostForm("password")
	password2 := ctx.PostForm("password2")
	username := ctx.DefaultPostForm("username", "mash")

	if password != password2 {
		panic("新密码和确认密码不一致")
	}

	var admin model.Admin
	model.ORM.Where("username = ?", username).Find(&admin)
	
	if !argon2idencoder.ComparePasswords(old_password, admin.Password) {
		panic("原始密码不正确")
	}

	password = argon2idencoder.EncodePassword(password)
	model.ORM.Model(&admin).Update("password", password)	

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "修改成功",
		"data": "",
	})
}