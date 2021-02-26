package admincontroller

import (
	"github.com/gin-gonic/gin"
	"api/model"
	"strconv"
	"log"
	"encoding/json"
	"fmt"
	"api/service/passwordencoder/argon2idencoder"
)

type adminForm struct {
	Username string		`form:"username" json:"username" binding:"required"`
	Password string		`form:"password"`
	RoleId int64		`form:"role_id"  binding:"required"`
	Descript string		`form:"descript"`
	IsEnabled int		`form:"is_enabled"`
}


func Index(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	log.Println(page)
	var adminModel model.Admin
	
	admins := adminModel.FindAll()
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": admins,
	})
}

func Add(ctx *gin.Context) {	
	/*dd, _ := ctx.GetRawData()
	log.Println(string(dd))
	var admin adminForm
	err := json.Unmarshal(dd, &admin)
	if err != nil {
		log.Println(err)
	}
	log.Println(admin)*/
	var admin adminForm
	if err := ctx.ShouldBind(&admin); err != nil {
		panic(err.Error())
	}
	b, _ := json.Marshal(&admin)
	
	var adminMap map[string]interface{}
	_ = json.Unmarshal(b, &adminMap)
	adminMap["Password"] = argon2idencoder.EncodePassword(adminMap["Password"].(string))
	
	var adminModel model.Admin
	adminModel.Insert(adminMap)
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg": "添加成功",
	})
	
}

func Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	var adminModel model.Admin
	admin := adminModel.Find(id)

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": admin,
	})
}

func EditSave(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	var adminModel model.Admin
	admin := adminModel.Find(id)
	
	var adminType adminForm
	if err = ctx.ShouldBind(&adminType); err != nil {
		panic(err.Error())
	}
	b, _ := json.Marshal(&adminType)
	var adminMap map[string]interface{}
	_ = json.Unmarshal(b, &adminMap)

	
	if adminMap["Password"] != "" {		
		adminMap["Password"] = argon2idencoder.EncodePassword(adminMap["Password"].(string))
	} else {
		adminMap["Password"] = admin.Password
	}

	admin.Update(adminMap, map[string]interface{}{"id":admin.Id})
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "编辑成功",
	})
}

func Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		panic(err)
	}

	var adminModel model.Admin
	adminModel.Delete(id)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "删除成功",
	})
}

func Status(ctx *gin.Context) {
	status := ctx.PostForm("status")
	id, err := strconv.Atoi(ctx.DefaultPostForm("id", "0"))
	if err != nil {
		panic(err)
	}

	var adminModel model.Admin
	admin := adminModel.Find(id)

	sql := fmt.Sprintf("update %s set is_enabled = ? where id = ?", adminModel.Tablename())
	stmt, err := model.Db.Prepare(sql)
	_, err = stmt.Exec(status, admin.Id)
	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
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

	var adminModel model.Admin
	var whereOrder model.OrderedMap
	whereOrder.Set("username", username)
	admin := adminModel.FindOneBy(whereOrder, model.OrderedMap{})


	if !argon2idencoder.ComparePasswords(old_password, admin.Password) {
		panic("原始密码不正确")
	}

	password = argon2idencoder.EncodePassword(password)

	sql := fmt.Sprintf("update %s set password = ? where id = ?", adminModel.Tablename())
	stmt, err := model.Db.Prepare(sql)
	_, err = stmt.Exec(password, admin.Id)
	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "修改成功",
		"data": "",
	})
}