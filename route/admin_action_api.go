package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/adminactionapicontroller"
)

func LoadAdminActionApi(router *gin.Engine) {
	router.GET("/api/admin/action/api", adminactionapicontroller.Api)
	router.POST("/api/admin/action/api/set", adminactionapicontroller.ApiSetAction)
}