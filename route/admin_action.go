package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/adminactioncontroller"
)

func LoadAdminAction(router *gin.Engine) {
	router.GET("/api/admin/action/index", adminactioncontroller.Index)
	router.POST("/api/admin/action/add", adminactioncontroller.Add)
	router.GET("/api/admin/action/edit/:id", adminactioncontroller.Edit)
	router.POST("/api/admin/action/edit/:id", adminactioncontroller.EditSave)
	router.POST("/api/admin/action/status", adminactioncontroller.Status)
	router.POST("/api/admin/action/delete", adminactioncontroller.Delete)
	router.POST("/api/admin/action/menu", adminactioncontroller.SetActionMenu)
	router.POST("/api/admin/action/collect", adminactioncontroller.CollectAction)
}