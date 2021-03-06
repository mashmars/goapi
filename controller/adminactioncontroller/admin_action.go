package adminactioncontroller

import (
	"api/model"
	"encoding/json"
	"strconv"
	"github.com/gin-gonic/gin"
	"math"
)

func Index(ctx *gin.Context) {
	var adminActions []model.AdminAction
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"));
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))	

	offset := (page - 1) * pageSize
	var count int64
	model.ORM.Model(&model.AdminAction{}).Where("is_sub_menu = 1").Count(&count)
	pages := math.Ceil(float64(count)/float64(pageSize))

	model.ORM.Where("is_sub_menu = 1").Select("admin_action.*, admin_menu.name as menu_name").Joins("left join admin_menu on admin_menu.id = admin_action.menu_id").Limit(pageSize).Offset(offset).Find(&adminActions)
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminActions,
		"totalPage": pages,
		"count": count,
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
	action_ids := ctx.PostFormArray("ids[]") //usestage html form.serialize 
	if len(action_ids) == 0 { //react js axios data
		data, _ := ctx.GetRawData()
		type PostData struct {
			Menu string
			Ids []int
		}
		var postData PostData
		if err := json.Unmarshal(data, &postData); err != nil {
			ctx.JSON(200, gin.H{
				"code": 1,
				"msg" : err,
				"data": "",
			})
			return
		}
		for _, id := range postData.Ids {
			action_ids = append(action_ids, strconv.Itoa(id))
		}		
		menu_id = postData.Menu
	}
	
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