package models

import "time"

type Tag struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Name      string     `gorm:"not null"`
	Posts     []Post     `gorm:"many2many:post_tags"`
}
