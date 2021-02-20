package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/adminrolecontroller"
)

func LoadAdminRole(router *gin.Engine) {
	router.GET("/api/admin/role/all", adminrolecontroller.All)
	router.GET("/api/admin/role/index", adminrolecontroller.Index)
	router.POST("/api/admin/role/add", adminrolecontroller.Add)
	router.GET("/api/admin/role/edit/:id", adminrolecontroller.Edit)
	router.POST("/api/admin/role/edit/:id", adminrolecontroller.EditSave)
	router.POST("/api/admin/role/status", adminrolecontroller.Status)
	router.POST("/api/admin/role/delete", adminrolecontroller.Delete)
}