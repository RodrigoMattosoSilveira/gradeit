package models

import (
	"gorm.io/gorm"
)

type PersonCreate struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type PersonUpdate struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type Person struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type PersonValidation struct {
	InvalidEmail     	bool `json:"invalid_email,omitempty"`
	EmailExists      	bool `json:"email_exists,omitempty"`
	InvalidPassword  	bool `json:"invalid_password,omitempty"`
	ParmIdInexistent 	bool `json:"parm_id_inexistent,omitempty"`
	InvalidParmId    	bool `json:"invalid_parm_id,omitempty"`
	PersonNotInDB    	bool `json:"person_not_in_db,omitempty"`
	NoUpdateAttributes	bool `json:"no_update-attributes,omitempty"`
}
