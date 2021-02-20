package adminactioncontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	var adminActions []model.AdminAction
	model.ORM.Where("is_sub_menu = 1").Find(&adminActions)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminActions,
	})
}

func Add(ctx *gin.Context) {
	var adminAction model.AdminAction
	if err := ctx.ShouldBind(&adminAction); err != nil {
		panic(err)
	}
	if result := model.ORM.Create(&adminAction); result.Error != nil {
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
	var adminAction model.AdminAction
	model.ORM.Find(&adminAction, id)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminAction,
	})
}

func EditSave(ctx *gin.Context) {
	id := ctx.Param("id")
	var adminAction model.AdminAction
	model.ORM.Find(&adminAction, id)

	if err := ctx.ShouldBind(&adminAction); err != nil {
		panic(err)
	}
	if result := model.ORM.Save(&adminAction); result.Error != nil {
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
	var adminAction model.AdminAction
	model.ORM.Find(&adminAction, id)
	model.ORM.Model(&adminAction).Update("is_enabled", status)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var adminAction model.AdminAction
	model.ORM.Find(&adminAction, id)
	model.ORM.Delete(&adminAction)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func SetActionMenu(ctx *gin.Context) {
	menu_id := ctx.PostForm("menu")
	action_ids := ctx.PostFormArray("ids[]")
	
	model.ORM.Model(&model.AdminAction{}).Where("id in ?", action_ids).Update("menu_id", menu_id)

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func CollectAction(ctx *gin.Context) {	
	var adminActions []model.AdminAction	
	if err := ctx.ShouldBind(&adminActions); err != nil {
		panic(err)
	}
	//存在不更新
	for _, adminAction := range adminActions {
		model.ORM.Where("router_name = ?", adminAction.RouterName).First(&adminAction)
		if adminAction.ID != 0 {
			continue
		} else {
			if result := model.ORM.Create(&adminAction); result.Error != nil {
				panic(result.Error)
			}
		}
	}
	/* or
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Printf("ctx.Request.body: %v", string(data))*/
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}

func CollectActionApi(ctx *gin.Context) {
	
}