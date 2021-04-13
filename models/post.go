package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string
	Content   string
	Published bool
	Slug      string
	AuthorID  uint `gorm:"foreignKey"`
	Author    Author
}
