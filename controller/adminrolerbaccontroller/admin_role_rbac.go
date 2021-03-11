package adminrolerbaccontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
)


//
func RbacInfo(ctx *gin.Context) {
	role_id := ctx.Param("id")
	var adminRole model.AdminRole
	model.ORM.Find(&adminRole, role_id)
	menus := menuInfo(role_id)
	roleMenus, roleActions, roleApis := menuRelation(role_id)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "",
		"data": map[string]interface{}{
			"menus": menus,
			"roleMenus": roleMenus,
			"roleActions": roleActions,
			"roleApis": roleApis,
		},
	})
}
//获取所有菜单
//获取菜单关联的action api
func menuInfo(role_id string) ([]map[string]interface{}) {
	var menus []map[string]interface{}
	model.ORM.Model(&model.AdminMenu{}).Select("id,name,sign").Where("is_enabled = 1").Find(&menus)//菜单
	
	for _, menu := range menus {
		var actions []map[string]interface{}
		model.ORM.Model(&model.AdminAction{}).Select("id, name").Where("menu_id = ? and is_enabled = 1", menu["id"]).Find(&actions)//菜单功能
		for _, action := range actions  {
			var apis []map[string]interface{}
			model.ORM.Model(&model.AdminActionApi{}).Select("id, method, path").Where("action_id = ? and is_enabled = 1", action["id"]).Find(&apis)//api
			action["apis"] = apis
		}
		//menu.Actions = actions
		menu["actions"] = actions
	}
	return menus
}

//已关联的信息
func menuRelation(role_id string) (roleMenus, roleActions, roleApis []int) {
	adminRoleMenus := []model.AdminRoleMenu{}
	model.ORM.Where("role_id = ?", role_id).Find(&adminRoleMenus)//所属菜单
	adminRoleActions := []model.AdminRoleAction{}
	model.ORM.Where("role_id = ?", role_id).Find(&adminRoleActions)//所属菜单子菜单
	adminRoleActionApis := []model.AdminRoleActionApi{}
	model.ORM.Where("role_id = ?", role_id).Find(&adminRoleActionApis)//所属菜单子菜单的api
	//var roleMenus, roleActions, roleApis []uint
	for _, row := range adminRoleMenus {
		roleMenus = append(roleMenus, row.MenuId)
	}
	for _, row := range adminRoleActions {
		roleActions = append(roleActions, row.ActionId)
	}
	for _, row := range adminRoleActionApis {
		roleApis = append(roleApis, row.ApiId)
	}
	return
}


//
func RbacSet(ctx *gin.Context) {
	role_id := ctx.Param("id")
	var adminRole model.AdminRole
	model.ORM.Find(&adminRole, role_id)

	//霸道一点处理 直接清空
	model.ORM.Where("role_id = ?", role_id).Delete(&model.AdminRoleMenu{})
	model.ORM.Where("role_id = ?", role_id).Delete(&model.AdminRoleAction{})
	model.ORM.Where("role_id = ?", role_id).Delete(&model.AdminRoleActionApi{})
	//加新的
	menus := ctx.PostFormArray("menu") //menu_action_auth menu_api_auth
	var adminMenu model.AdminMenu
	roleMenus := []model.AdminRoleMenu{}
	roleActions := []model.AdminRoleAction{}
	roleActionApis := []model.AdminRoleActionApi{}
	for _, menu_id := range menus {
		model.ORM.Find(&adminMenu, menu_id)
		actions := ctx.PostFormArray("menu_action_" + adminMenu.Sign)		
		//添加
		roleMenu := model.AdminRoleMenu{}
		roleMenu.RoleId, _ = strconv.Atoi(role_id)
		roleMenu.MenuId, _ = strconv.Atoi(menu_id)
		roleMenus = append(roleMenus, roleMenu)
		for _, action_id := range actions {
			roleAction := model.AdminRoleAction{}
			roleAction.RoleId, _ = strconv.Atoi(role_id)
			roleAction.ActionId, _ = strconv.Atoi(action_id)
			roleActions = append(roleActions, roleAction)
			apis := ctx.PostFormArray("menu_api_" + adminMenu.Sign + "_" + action_id)			
			for _, api_id := range apis {
				roleActionApi := model.AdminRoleActionApi{}
				roleActionApi.RoleId, _ = strconv.Atoi(role_id)
				roleActionApi.ActionId, _ = strconv.Atoi(action_id)
				roleActionApi.ApiId, _ = strconv.Atoi(api_id)
				roleActionApis = append(roleActionApis, roleActionApi)
			}			
		}
	}

	model.ORM.Create(&roleMenus)
	model.ORM.Create(&roleActions)
	model.ORM.Create(&roleActionApis)
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": "",
	})
}


//
func RbacSetReact(ctx *gin.Context) {
	role_id := ctx.Param("id")
	var adminRole model.AdminRole
	model.ORM.Find(&adminRole, role_id)

	//霸道一点处理 直接清空
	model.ORM.Where("role_id = ?", role_id).Delete(&model.AdminRoleMenu{})
	model.ORM.Where("role_id = ?", role_id).Delete(&model.AdminRoleAction{})
	model.ORM.Where("role_id = ?", role_id).Delete(&model.AdminRoleActionApi{})
	//加新的
	type PostData struct {
		Menus []int
		Actions []int
		Apis []int
	}

	rawData,_ := ctx.GetRawData()
	var postData PostData
	if err := json.Unmarshal(rawData, &postData); err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg" : err,
			"data": "",
		})
		return
	}	

	roleMenu := model.AdminRoleMenu{}
	roleMenus := []model.AdminRoleMenu{}

	roleActions := []model.AdminRoleAction{}
	roleActionApis := []model.AdminRoleActionApi{}

	roleAction := model.AdminRoleAction{}
	roleActionApi := model.AdminRoleActionApi{}

	for _, menu_id := range postData.Menus {
		roleMenu.RoleId, _ = strconv.Atoi(role_id)
		roleMenu.MenuId = menu_id
		roleMenus = append(roleMenus, roleMenu)
	}

	for _, action_id := range postData.Actions {
		roleAction.RoleId, _ = strconv.Atoi(role_id)
		roleAction.ActionId = action_id
		roleActions = append(roleActions, roleAction)
	}

	for _, api_id := range postData.Apis {
		roleActionApi.RoleId, _ = strconv.Atoi(role_id)
		actionApi := model.AdminActionApi{}
		model.ORM.Find(&actionApi, api_id)
		roleActionApi.ActionId = actionApi.ActionId
		roleActionApi.ApiId = api_id
		roleActionApis = append(roleActionApis, roleActionApi)
	}

	model.ORM.Create(&roleMenus)
	model.ORM.Create(&roleActions)
	model.ORM.Create(&roleActionApis)
	
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": "",
	})
}