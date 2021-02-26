package model

import (
	"gorm.io/gorm"
	"errors"
)


type AdminRoleMenu struct {
	ID int			`json:"id" gorm:"type:int(10) unsigned auto_increment;"`
	RoleId int		`json:"role_id" form:"role_id" gorm:"type:int(10)"`
	MenuId int		`json:"menu_id" form:"menu_id" gorm:"type:int(10)"`
}

type AdminRoleAction struct {
	ID int			`json:"id" gorm:"type:int(10) unsigned auto_increment;"`
	RoleId int		`json:"role_id" form:"role_id" gorm:"type:int(10)"`
	ActionId int	`json:"action_id" form:"action_id" gorm:"type:int(10)"`
}

type AdminRoleActionApi struct {
	ID int			`json:"id" gorm:"type:int(10) unsigned auto_increment;"`
	RoleId int		`json:"role_id" form:"role_id" gorm:"type:int(10)"`
	ActionId int	`json:"action_id" form:"action_id" gorm:"type:int(10)"`
	ApiId int		`json:"api_id" form:"api_id" gorm:"type:int(10)"`
}

func (adminRoleMenu *AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}

func (adminRoleAction *AdminRoleAction) TableName() string {
	return "admin_role_action"
}

func (adminRoleActionApi *AdminRoleActionApi) TableName() string {
	return "admin_role_action_api"
}


func (adminRoleMenu *AdminRoleMenu) BeforeDelete(tx *gorm.DB) (err error) {
	if adminRoleMenu.RoleId == 1 {
		panic(errors.New("超级管理员权限不能删除"))
	}
	return
}

func (adminRoleAction *AdminRoleAction) BeforeDelete(tx *gorm.DB) (err error) {
	if adminRoleAction.RoleId == 1 {
		panic(errors.New("超级管理员权限不能删除"))
	}
	return
}

func (adminRoleActionApi *AdminRoleActionApi) BeforeDelete(tx *gorm.DB) (err error) {
	if adminRoleActionApi.RoleId == 1 {
		panic(errors.New("超级管理员权限不能删除"))
	}
	return
}