package route

import (
	"github.com/gin-gonic/gin"
	"api/controller/admincontroller"
)

func LoadAdmin(router *gin.Engine) {
	admin := router.Group("/api/admin")
	{
		admin.GET("/index", admincontroller.Index)	
		admin.POST("/add", admincontroller.Add)
		admin.GET("/edit/:id", admincontroller.Edit)
		admin.POST("/edit/:id", admincontroller.EditSave)
		admin.POST("/delete", admincontroller.Delete)
		admin.POST("/status", admincontroller.Status)
		admin.POST("/password", admincontroller.Password)
	}
}