package assignment

import (
	"regexp"
	"strconv"
	"time"

	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

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
// TODO Figure out a way to unit test it
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
// TODO Figure out a way to unit test it
func PersonInDB(id uint64) bool {
	var person models.Person
	result := configs.DB.First(&person, id)
	return  result.Error == nil
}

// Ensure the int64 is the ID of an assigment in the databse
// 
// Input:   int64    
// 
// Output:  true if the id is that of an assignment in the database, false otherwise
//
// TODO Figure out a way to unit test it
func AssignmentInDB(id uint64) bool {
	var assignment models.Assignment
	result := configs.DB.First(&assignment, id)
	return  result.Error == nil
}

// Bind the HTTP request id parameter
// 
// Input:   *gin.Context    
// 
// Output:  (true, id) able to bind it,  (false, 0) otherwise
//
// TODO Figure out a way to unit test it
// func ParseIdParm(ctx *gin.Context) (bool, string) {
// 	idParm := ctx.Param("id")
// 	if idParm == "" {
// 		return false, ""
// 	} 
// 	return true, idParm
// }

// Ensure the string is a valid uint64
// 
// Input:   string    
// 
// Output:  (id, true) if thestring is a valid uint64,  (0, false) otherwise
//
func ValidIdParm(idParm string) (uint64, bool) {
	if idParm == "" {
		return 0, false
	}
	id, err := strconv.ParseUint(idParm, 10, 32)
	if err != nil {
		return 0, false
	}
	return id, true
}

// Ensure the time, due date, is beyond now
// 
// Input:   time.time    
// 
// Output:  true if time, due date, is beyond now,  false otherwise
//
func DueNotOld(due time.Time) bool {
	//  Reduce everything to the midnight UTC prior the time
	timeNowMidNight := time.Now()
	timeNowMidNight = time.Date(timeNowMidNight.Year(), timeNowMidNight.Month(), timeNowMidNight.Day(), 0, 0, 0, 0, time.UTC)
	dueMidNight := time.Date(due.Year(), due.Month(), due.Day(), 0, 0, 0, 0, time.UTC)
	return !dueMidNight.Before(timeNowMidNight)

}