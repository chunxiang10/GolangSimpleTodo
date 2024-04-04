package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func init() {
	var err error
	path := "data/todo.db"
	Db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "todo_",
		},
	})
	if err != nil {
		panic("无法连接数据库" + err.Error())

	}
	sqlDB, err := Db.DB()
	if err != nil {
		panic("初始化数据库失败" + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func DbInit() *gorm.DB {
	return Db
}
