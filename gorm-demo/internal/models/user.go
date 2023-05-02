package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Username  string     `gorm:"not null"`
	Email     string     `gorm:"not null"`
	Password  string     `gorm:"not null"`
	Posts     []Post     `gorm:"foreignKey:UserID"`
	Comments  []Comment  `gorm:"foreignKey:UserID"`
}
