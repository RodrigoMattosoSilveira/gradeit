package controllers

import (
	"regexp"
	"strconv"

	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	"github.com/gin-gonic/gin"
)

// Ensure the string is a valid email address
//
// Input:   string
//
// Output:  true if a valid email address, false otherwise
//
func GetIdParm(email string) bool {
	// regex pattern for email ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// RegexPattern := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	// TODO find a more robust validation REGEX 
	RegexPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return RegexPattern.MatchString(email)

}

// Ensure the string is a valid email address
//
// Input:   string
//
// Output:  true if a valid email address, false otherwise
//
func ValidEmail(email string) bool {
	// regex pattern for email ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// RegexPattern := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	// TODO find a more robust validation REGEX 
	RegexPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return RegexPattern.MatchString(email)

}

// Ensure that the string is not an email of person in the database
// 
// Input:   string    
// 
// Output:  true if a record with this email does not exist, false otherwise
// 
func UniqueEmail(email string) bool {
	var person models.Person
	result := configs.DB.Where("email = ?", email).First(&person)
	return  result.Error != nil
}

// Ensure the string is a valid password
// 
// Input:   string    
// 
// Output:  true if a valid password, false otherwise
//
func ValidPassword(password string) bool {
	// TODO add a robust REGEX to validate this
	return  len(password)>5
}

// Ensure the int64 is the ID of a person in the databse
// 
// Input:   int64    
// 
// Output:  true if the id is that of a person in the database, false otherwise
//
func PersonInDB(id uint64) bool {
	var person models.Person
	result := configs.DB.First(&person, id)
	return  result.Error == nil
}

// Bind the HTTP request id parameter
// 
// Input:   *gin.Context    
// 
// Output:  (true, id) able to bind it,  (false, 0) otherwise
//
func ParseIdParm(ctx *gin.Context) (bool, string) {
	idParm := ctx.Param("id")
	if idParm == "" {
		return false, ""
	} 
	return true, idParm
}

// Ensure the string is a valid uint64
// 
// Input:   string    
// 
// Output:  (true, id) if thestring is a valid uint64,  (false, 0) otherwise
//
func ValidIdParm(idParm string) (bool, uint64) {
	goodIdParm :=true
	id, err := strconv.ParseUint(idParm, 10, 32)
	if err != nil {
		goodIdParm = false 
		id = 0
	}
	return goodIdParm, id
}
