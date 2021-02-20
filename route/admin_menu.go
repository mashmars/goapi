package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/adminmenucontroller"
)

func LoadAdminMenu(router *gin.Engine) {
	router.GET("/api/admin/menu/all", adminmenucontroller.All)
	router.GET("/api/admin/menu/index", adminmenucontroller.Index)
	router.POST("/api/admin/menu/add", adminmenucontroller.Add)
	router.GET("/api/admin/menu/edit/:id", adminmenucontroller.Edit)
	router.POST("/api/admin/menu/edit/:id", adminmenucontroller.EditSave)
	router.POST("/api/admin/menu/status", adminmenucontroller.Status)
	router.POST("/api/admin/menu/delete", adminmenucontroller.Delete)
	router.GET("/api/admin/menu/action/:id", adminmenucontroller.MenuAction)
}