package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
	Nick      string
	Avatar    string
	Email     string
	Password  string
}
