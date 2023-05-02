package models

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Body      string     `gorm:"not null"`
	UserID    uint       `gorm:"not null"`
	User      User
	PostID    uint `gorm:"not null"`
	Post      Post
}
