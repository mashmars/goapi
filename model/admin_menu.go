package model

import (
	"gorm.io/gorm"
)

type AdminMenu struct {
	ID int				`json:"id" form:"id" gorm:"type:int(10) unsigned AUTO_INCREMENT not null;"`
	Name string			`json:"name" form:"name" binding:"required" gorm:"size:100"`
	Sign string			`json:"sign" form:"sign" binding:"required" gorm:"uniqueindex;size:20;comment:唯一标识"`
	SortedBy int		`json:"sorted_by" form:"sorted_by" gorm:"default:100;type:smallint"`
	Icon string			`json:"icon" form:"icon" gorm:"size:30;default:fab fa-symfony;"`
	IsEnabled int		`json:"is_enabled" form:"is_enabled" gorm:"type:tinyint"`	
}

func (adminMenu *AdminMenu) TableName() string {
	return "admin_menu"
}


func (adminMenu *AdminMenu) AfterFind(tx *gorm.DB) (err error) {
	
	if adminMenu.ID == 0 {
		panic("请求信息不正确")
	}
	return
}
