package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/adminrolerbaccontroller"
)

func LoadAdminRoleRbac(router *gin.Engine) {
	router.GET("/api/admin/role/rbac/:id", adminrolerbaccontroller.RbacInfo)
	router.POST("/api/admin/role/rbac/set/:id", adminrolerbaccontroller.RbacSet)
	router.POST("/api/admin/role/rbac/set1/:id", adminrolerbaccontroller.RbacSetReact)
}