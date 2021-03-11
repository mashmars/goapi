package securitycontroller

import (
	"github.com/gin-gonic/gin"
	"api/model"
	"api/service/passwordencoder/argon2idencoder"
	"api/service/jwtmanager"
	"encoding/json"
)

func Login(ctx *gin.Context) {
	var adminModel model.Admin
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
	//find 根据主键返回row
	where := model.OrderedMap{}
	where.Set("username", username)
	admin := adminModel.FindOneBy(where, model.OrderedMap{})

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
	
	ctx.JSON(200, gin.H{
		"code": 0, 
		"msg" : "登陆成功",
		"data": map[string]interface{}{
			"token": token,
			"username": username,
		},
	})
}