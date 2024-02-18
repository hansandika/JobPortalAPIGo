package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	RoleId    int    `gorm:"default:1"`
	RoleName  string `gorm:"default:'user'"`
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
