package data

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title       string
	Description string
	Body        string
	CreateAt    time.Time
	UpdateAt    time.Time
}
