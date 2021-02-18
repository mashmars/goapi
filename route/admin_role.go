package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/adminrolecontroller"
)

func LoadAdminRole(router *gin.Engine) {
	router.GET("/api/admin/role/all", adminrolecontroller.All)
}