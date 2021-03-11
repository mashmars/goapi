package adminactionapicontroller

import (
	"api/model"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"strconv"
)

func Api(ctx *gin.Context) {
	var adminActionApis []model.AdminActionApi
	model.ORM.Where("is_enabled = 1").Order("controller_action asc").Find(&adminActionApis)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "success",
		"data": adminActionApis,
	})
}


func ApiSetAction(ctx *gin.Context) {
	id := ctx.PostForm("action_id") //action id
	ids := ctx.PostFormArray("ids[]") //api id

	if len(ids) == 0 { //react js axios data
		data, _ := ctx.GetRawData()
		type PostData struct {
			ActionId string 	`json:"action_id"`
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
			ids = append(ids, strconv.Itoa(id))
		}		
		id = postData.ActionId
	}
	
	model.ORM.Model(&model.AdminActionApi{}).Where("action_id = ?", id).Update("action_id", 0)
	model.ORM.Model(&model.AdminActionApi{}).Where("id in (?)", ids).Update("action_id", id) 
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg" : "操作成功",
		"data": "",
	})
}