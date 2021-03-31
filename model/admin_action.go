package model

import (
	"gorm.io/gorm"
	"errors"
)

type AdminAction struct {
	ID int 						`json:"id" form:"id" gorm:"type:int(10) unsigned auto_increment not null"`
	MenuId int					`json:"menu_id" form:"menu_id" gorm:"type:int(10)"`
	Name string					`json:"name" form:"name" gorm:"size:30;"`
	RouterName string			`json:"router_name" form:"router_name" gorm:"uniqueIndex;size:100;"`
	RouterShortName string		`json:"router_short_name" form:"router_short_name" gorm:"size:100;"`
	IsSubMenu int				`json:"is_sub_menu" form:"is_sub_menu" gorm:"type:tinyint;default:0;"`
	SortedBy int				`json:"sorted_by" form:"sorted_by" gorm:"default:100;type:smallint"`
	Icon string					`json:"icon" form:"icon" gorm:"size:100;default:far fa-file;"`
	ControllerAction string		`json:"controller_action" form:"controller_action" gorm:"size:150;"`
	IsEnabled int				`json:"is_enabled" form:"is_enabled" gorm:"type:tinyint;"`
	MenuName string				`json:"menu_name" form:"menu_name" gorm:"->"`
}


func (adminAction *AdminAction) TableName() string {
	return "admin_action"
}

func (adminAction *AdminAction) AfterFind(tx *gorm.DB) (err error) {
	if adminAction.ID == 0 {
		panic(errors.New("请求信息不正确"))
	}
	return
}