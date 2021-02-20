package model

type AdminActionApi struct {
	ID int 						`json:"id" form:"id" gorm:"type:int(10) unsigned auto_increment not null"`
	ActionId int				`json:"action_id" form:"action_id" gorm:"type:int(10)"`
	Name string					`json:"name" form:"name" gorm:"size:30;"`
	Method string				`json:"method" form:"method" gorm:"size:10;"`
	Path string					`json:"path" form:"path" gorm:"size:100;"`
	SortedBy int				`json:"sorted_by" form:"sorted_by" gorm:"default:100;type:smallint"`
	ControllerAction string		`json:"controller_action" form:"controller_action" gorm:"size:150;"`
	IsEnabled int				`json:"is_enabled" form:"is_enabled" gorm:"type:tinyint;"`
}


func (adminActionApi *AdminActionApi) TableName() string {
	return "admin_action_api"
}