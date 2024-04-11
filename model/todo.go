package model

import (
	"simple-todo/db"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int64          `json:"id" gorm:"column:id;primarykey"`
	Content   string         `json:"content" gorm:"column:content"`
	Completed bool           `json:"completed" gorm:"column:completed;default:false"`
	Display   bool           `json:"display" gorm:"column:display;default:true"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}

func init() {
	d := db.DbInit()
	d.AutoMigrate(&Todo{})
}
