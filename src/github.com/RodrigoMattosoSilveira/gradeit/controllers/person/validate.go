package controllers

import (
	"fmt"
	"regexp"

	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

func validEmail(email string) bool {
	// regex pattern for email ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// RegexPattern := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	// TODO find a more robust validation REGEX 
	RegexPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return RegexPattern.MatchString(email)

}


func validPassword(password string) bool {
	return  len(password)>0
}

func ValidatePersonCreate(person models.Person) (bool, []string) {
	validPerson := true
	errors := make([]string, 0)

	if !validEmail(person.Email) {
		validPerson = false
		errors = append(errors, fmt.Sprintf("Person %d: invalid email = %s", person.ID, person.Email))
	}

	if !validPassword(person.Email) {
		validPerson = false
		errors = append(errors, fmt.Sprintf("Person %d: invalid password = %s", person.ID, person.Password))
	}

	return validPerson, errors
}


func ValidatePersonUpdate(person models.Person) (bool, []string) {
	validPerson := true
	errors := make([]string, 0)

	
	// if person.Email !="" && !validEmail(person.Email) {
	// 	validPerson = false
	// 	errors = append(errors, fmt.Sprintf("Person %d: invalid email = %s", person.ID, person.Email))
	// }

	if person.Password !="" && !validPassword(person.Email) {
		validPerson = false
		errors = append(errors, fmt.Sprintf("Person %d: invalid password = %s", person.ID, person.Password))
	}

	return validPerson, errors
}
