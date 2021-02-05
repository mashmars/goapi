package admincontroller

import (
	"github.com/gin-gonic/gin"
	"api/model"
	"strconv"
	"log"
	"encoding/json"
)

type adminForm struct {
	Username string		`form:"username" json:"username" binding:"required"`
	Password string		`form:"password" binding:"required"`
	RoleId int64		`form:"role_id"  binding:"required"`
	Descript string		`form:"descirpt"`
	IsEnabled int		`form:"is_enabled"`
}


func Index(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	log.Println(page)
	var adminModel model.Admin
	admins, _ := adminModel.FindAll()
	
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
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg" : err.Error(),
		})
		return
	}
	b, _ := json.Marshal(&admin)
	
	var adminMap map[string]interface{}
	_ = json.Unmarshal(b, &adminMap)
	
	var adminModel model.Admin
	_, err := adminModel.Insert(adminMap)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg": err,
		})
		return
	}
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
	admin, err := adminModel.Find(id)
	if err != nil {
		panic(err)
	}

	if admin == nil {

	}

	var adminType adminForm
	if err = ctx.ShouldBind(&adminType); err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg" : err.Error(),
		})
		return
	}
	b, _ := json.Marshal(&adminType)
	var adminMap map[string]interface{}
	_ = json.Unmarshal(b, &adminMap)
	_, err = admin.Update(adminMap, map[string]interface{}{"id":admin.Id})
	if err != nil {
		panic(err)
	}
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