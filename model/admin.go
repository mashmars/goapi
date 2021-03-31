package model

import (
	"time"
)

type Admin struct {
	Id int					`json:"id" form:"id" gorm:"type:int(10) unsigned auto_increment not null"`
	Username string			`json:"username" form:"username" gorm:"size:30;"`
	Password string 		`json:"password" form:"password" gorm:"size:150;"`
	RoleId int				`json:"role_id" form:"role_id" gorm:"type:int(10)"`
	Descript string			`json:"descript" form:"descript" gorm:"size:150;"`
	IsEnabled int			`json:"is_enabled" form:"is_enabled" gorm:"type:tinyint"`
	LastLoginIp string		`json:"last_login_ip" form:"name" gorm:"size:30;"`
	CreatedAt time.Time 	`json:"created_at"`
	LastLoginAt time.Time	`json:"last_login_at"`
	RoleName string			`json:"role_name" form:"role_name" gorm:"->"`
}

func (admin *Admin) TableName() string {
	return "admin"
}