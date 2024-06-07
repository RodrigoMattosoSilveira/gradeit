package models

import (
	"time"

	"gorm.io/gorm"
)

type AssignmentCreate struct {
	PersonId	uint64		`form:"person_id" json:"person_id" binding:"required"`
	Description	string		`form:"description" json:"description" binding:"required"`
	Due 		time.Time	`form:"due" json:"due" binding:"required"`
} 
type AssignmentUpdate struct {
	PersonId	uint64		`form:"person_id" json:"person_id"`
	Description	string		`form:"description" json:"description"`
	Due 		time.Time	`form:"due" json:"due"`
}

type Assignment struct {
	gorm.Model
	ID        	uint64		`gorm:"primaryKey"`
	PersonId	uint64       `gorm:"person_id"`
	Description	string
	Due 		time.Time
}