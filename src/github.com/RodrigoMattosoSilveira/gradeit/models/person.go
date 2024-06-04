package models

import (
	"gorm.io/gorm"
)

type PersonCreate struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}
type PersonUpdate struct {
	Name      string `form:"name" json:"name"`
	Email     string `form:"email" json:"email"`
	Password  string `form:"password" json:"password"`
}

type Person struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
}
