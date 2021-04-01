package securitycontroller

import (
	"github.com/gin-gonic/gin"
	"api/model"
	"api/service/passwordencoder/argon2idencoder"
	"api/service/jwtmanager"
	"encoding/json"
	"log"
)

func Login(ctx *gin.Context) {
	var admin model.Admin
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" { //兼容 react
		rawData, _ := ctx.GetRawData()
		type PostData struct {
			Username string `json:"username"`
			Password string	`json:"password"`
		}
		var postData PostData
		if err := json.Unmarshal(rawData, &postData); err != nil {
			ctx.JSON(200, gin.H{
				"code": 1,
				"msg" : err,
				"data": "",
			})
			return
		}
		username = postData.Username
		password = postData.Password
	}
 
	if username == "" {
		ctx.JSON(200, gin.H{
			"code": 1, 
			"msg" : "登陆失败， 用户名或密码错误",
			"data": "",
		})
		return
	}
	
	
	model.ORM.Where("username = ?", username).Find(&admin)

	if match := argon2idencoder.ComparePasswords(password, admin.Password); !match {
		ctx.JSON(200, gin.H{
			"code": 1, 
			"msg" : "登陆失败， 用户名或密码错误",
			"data": "",
		})
		return
	}

	tokenMap := map[string]interface{}{
		"username": username,
	}
	
	token, err := jwtmanager.CreateJwtToken(tokenMap)
	if err != nil {
		panic(err)
	}

	//获取已经授权的菜单等信息
	menus, actions, apis := getAuthorize(admin)
	log.Println(actions)
	ctx.JSON(200, gin.H{
		"code": 0, 
		"msg" : "登陆成功",
		"data": map[string]interface{}{
			"token": map[string]interface{}{
				"token": token,
				"username": username,
			},
			"authorize": map[string]interface{}{
				"roleMenus": menus,
				"roleActions": actions,
				"roleApis": apis,
			},
		},
	})
}


func getAuthorize(admin model.Admin) (menus, actions, apis []map[string]interface{}) {	
	if admin.Id == 1 {
		//默认超级管理员
		model.ORM.Model(&model.AdminMenu{}).Select("admin_menu.sign").Where("is_enabled = 1").Find(&menus)
		model.ORM.Model(&model.AdminAction{}).Select("admin_action.router_name").Where("is_enabled = 1").Find(&actions)
		model.ORM.Model(&model.AdminActionApi{}).Select("admin_action_api.method, admin_action_api.path").Where("is_enabled = 1").Find(&apis)
	} else {
		model.ORM.Model(&model.AdminMenu{}).Joins("left join admin_role_menu on admin_role_menu.menu_id = admin_menu.id").
		Select("admin_menu.sign").Where("admin_menu.is_enabled = 1 and admin_role_menu.role_id = ?", admin.RoleId).Find(&menus)
		model.ORM.Model(&model.AdminAction{}).Joins("left join admin_role_action on admin_role_action.action_id = admin_action.id").
		Select("admin_action.router_name").Where("admin_action.is_enabled = 1 and admin_role_action.role_id = ?", admin.RoleId).Find(&actions)
		model.ORM.Model(&model.AdminActionApi{}).Joins("left join admin_role_action_api on admin_role_action_api.api_id = admin_action_api.id").
		Select("admin_action_api.method, admin_action_api.path").Where("admin_action_api.is_enabled = 1 and admin_role_action_api.role_id = ?", admin.RoleId).Find(&apis)
	}
	return
}