package model

import (
	"gorm.io/gorm"
	"errors"
)

type AdminRole struct {
	ID uint			`json:"id"`
	Name string		`json:"name" form:"name" binding:"required"`
	IsEnabled int	`json:"is_enabled" form:"is_enabled"`
}

func (adminRole *AdminRole) TableName() string {
	return "admin_role"
}


func (adminRole *AdminRole) BeforeDelete(tx *gorm.DB) (err error) {
	if adminRole.ID == 1 || adminRole.Name == "超级管理员" {
		panic(errors.New("超级管理员角色不能删除"))
	}
	return
}