package adminmenucontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"strconv"
	"math"
)

func Index(ctx *gin.Context) {
	var adminMenus []model.AdminMenu
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"));
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))	

	offset := (page - 1) * pageSize
	var count int64
	model.ORM.Model(&model.AdminMenu{}).Count(&count)
	pages := math.Ceil(float64(count)/float64(pageSize))

	model.ORM.Limit(pageSize).Offset(offset).Find(&adminMenus)
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminMenus,
		"totalPage": pages,
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

func Check(ctx *gin.Context) {
	routes, _ := ctx.GetRawData()

	type SubAction struct {
		Name string 
		Path string 		
	}

	type Action struct {
		Name string 
		Path string 
		Prefix string 
		SubMenus []SubAction
	}

	type Menu struct {
		MenuName string 
		Sign string 
		SubRoutes []Action
	}

	var menus []Menu
    if err := json.Unmarshal(routes, &menus); err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg" : err,
			"data": "",
		})
		return
	}

	var adminMenus []model.AdminMenu
	var adminActions []model.AdminAction
	for _, menu := range menus {
		for _, action := range menu.SubRoutes {
			for _, subAction := range action.SubMenus {
				result := map[string]interface{}{}
				model.ORM.Model(&model.AdminAction{}).Where("router_name = ?", subAction.Path).First(&result)
				if _, ok := result["id"]; ok {
					continue;
				}
				adminActions = append(adminActions, model.AdminAction{
					Name: subAction.Name,
					RouterName: subAction.Path,
					RouterShortName: action.Prefix,
					IsEnabled: 1,
				})
			}
			result := map[string]interface{}{}
			model.ORM.Model(&model.AdminAction{}).Where("router_name = ?", action.Path).First(&result)
			if _, ok := result["id"]; ok {
				continue;
			}
			adminActions = append(adminActions, model.AdminAction{
				Name: action.Name,
				RouterName: action.Path,
				RouterShortName: action.Prefix,
				IsSubMenu: 1,
				IsEnabled: 1,
			})
		}
		result := map[string]interface{}{}
		model.ORM.Model(&model.AdminMenu{}).Where("sign = ?", menu.Sign).First(&result)
	
		if _, ok := result["id"]; ok {
			continue;
		}
		
		adminMenus = append(adminMenus, model.AdminMenu{
			Name: menu.MenuName,
			Sign: menu.Sign,
			IsEnabled: 1,
		})
	}

	model.ORM.Create(&adminMenus)
	model.ORM.Create(&adminActions)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})	
}