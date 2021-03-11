package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ORM *gorm.DB

func init() {
	//dsn := "root:mash@tcp(127.0.0.1:3306)/sfadmin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:mash@tcp(127.0.0.1:3306)/reactadmin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	ORM, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
		//DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, //取消默认事务 不自动开启
		DisableForeignKeyConstraintWhenMigrating: true, //取消创建外键约束
	})	
	/* or
	sqlDB, err := sql.Open("mysql", "root:mash@tcp(127.0.0.1:3306)/sfadmin?charset=utf8mb4&parseTime=True&loc=Local")
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})*/
	
	if err != nil {
		panic("failed to connect database")
	}

	//
	migrate()	
}


func migrate() {
	//ORM.AutoMigrate(&AdminMenu{})
	//ORM.AutoMigrate(&AdminAction{})
	//ORM.AutoMigrate(&AdminActionApi{})
	//ORM.AutoMigrate(&AdminRoleMenu{}, &AdminRoleAction{}, &AdminRoleActionApi{})
}