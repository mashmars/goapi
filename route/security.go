package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/securitycontroller"
)

func LoadSecurity(router *gin.Engine) {
	router.POST("/api/admin/login", securitycontroller.Login)
}