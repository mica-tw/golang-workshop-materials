package models

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Title     string     `gorm:"not null"`
	Body      string     `gorm:"not null"`
	UserID    uint       `gorm:"not null"`
	User      User
	Comments  []Comment `gorm:"foreignKey:PostID"`
	Tags      []Tag     `gorm:"many2many:post_tags"`
}
