package models

type Subscriber struct {
	ID       int64
	email    string
	AuthorID uint `gorm:"foreignKey"`
	Author   Author
}
