package model

import (
	"simple-todo/db"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id" gorm:"column:id;primarykey"`
	Uid       string         `json:"uid" gorm:"column:uid;unique;size:4"`
	Username  string         `json:"username" gorm:"column:username"`
	Password  string         `json:"password" gorm:"column:password"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}

func init() {
	d := db.DbInit()
	// var user = &User{Uid: "admin", Password: utils.EncryMd5("123456")}
	// d.Create(&user)
	d.AutoMigrate(&User{})
}
