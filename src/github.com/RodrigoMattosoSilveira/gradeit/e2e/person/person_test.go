package main

import (
	// "net/http"
	// "encoding/json"
	// "fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	"github.com/RodrigoMattosoSilveira/gradeit/routes"

	"github.com/Jeffail/gabs"
)

var albertEinstein = models.Person {
	Name: "Albert Einstein",
	Email:  strings.ToLower("Albert.Einstein@gmail.com"),
	Password: "Albert.Einstein123",
}

var francisBacon = models.Person {
	Name: "Francis Bacon",
	Email:  strings.ToLower("Francis.Bacon@gmail.com"),
	Password: "Francis.Bacon123",
}

// var antoineLavoisier = models.PersonCreate {
// 	Name: "Antoine Lavoisier",
// 	Email:  strings.ToLower("Antoine.Lavoisier@gmail.com"),
// 	Password: "Antoine.Lavoisier123",
// }

var neilsBohr = models.Person {
	Name: "Neils Bohr",
	Email:  strings.ToLower("Neils.Bohrn@gmail.com"),
	Password: "Neils.Bohr123",
}

//  Inspired by
// // https://blog.marcnuri.com/go-testing-gin-gonic-with-httptest
// 
func TestPerson(t *testing.T) {
	// setup logic
	configs.Config()
	router := configs.GetRouter()
	routes.RoutesPerson(router)
	setupPersonTable()
	t.Run("CreateInvalidEmail", func(t *testing.T) { 

		// personString := `{"Name": "Neils Bohr", "Email": "Neils.Bohr@gmail", "Password": "Neils.Bohr123abc"}`
		// req, _ := http.NewRequest("POST", "/person", strings.NewReader(personString))

		// Passing a structure to an HTTP call is tricky!	
		// neilsBohrWork := neilsBohr
		// neilsBohrWork.Email =  strings.ToLower("Neils.Bohrn@gm")
		// reqBodyBytes := new(bytes.Buffer)
		// json.NewEncoder(reqBodyBytes).Encode(neilsBohrWork)
		reqBodyBytes := getPersonWithInvalidAttributeSerialized(neilsBohr, "email", "bad@email")
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, req)
		if recorder.Code != 422 {
			t.Error("Expected 422, got ", recorder.Code)
		}		
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		invalidEmail, ok := jsonParsed.Path("error.invalid_email").Data().(bool) 
		if !ok {
			panic("Expected E2E CreateInvalidEmail parsing error")
		} 
	 	if !invalidEmail {
			t.Error("Expected Person email to be invalid, got ", invalidEmail)
		} 
	})
	t.Run("CreateExistingEmail", func(t *testing.T) { 

		// Passing a structure to an HTTP call is tricky!	
		// neilsBohrWork := neilsBohr
		// neilsBohrWork.Email =  strings.ToLower("Albert.Einstein@gmail.com")
		// reqBodyBytes := new(bytes.Buffer)
		// json.NewEncoder(reqBodyBytes).Encode(neilsBohrWork)
		reqBodyBytes := getPersonWithInvalidAttributeSerialized(neilsBohr, "email", "Albert.Einstein@gmail.com")
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, req)
		if recorder.Code != 422 {
			t.Error("Expected 422, got ", recorder.Code)
		}
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		emailExists, ok := jsonParsed.Path("error.email_exists").Data().(bool) 
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		} 
		if !emailExists {
			t.Error("Expected Person email to exist", emailExists)
		}
	 })
	t.Run("CreateInvalidPassword", func(t *testing.T) { 

		// Passing a structure to an HTTP call is tricky!	
		// neilsBohrWork := neilsBohr
		// neilsBohrWork.Password = "Nei"
		// reqBodyBytes := new(bytes.Buffer)
		// json.NewEncoder(reqBodyBytes).Encode(neilsBohrWork)
		reqBodyBytes := getPersonWithInvalidAttributeSerialized(neilsBohr, "password", "Nei")
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, req)
		if recorder.Code != 422 {
			t.Error("Expected 422, got ", recorder.Code)
		  }			  
		  jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		  if err != nil {
			  panic(err)
		  }
		  invalidPassword, ok := jsonParsed.Path("error.invalid_password").Data().(bool) 
		  if !ok {
			  panic("Expected E2E CreateInvalidPassword parsing error")
		  } 
		   if !invalidPassword {
			  t.Error("Expected Person password to be invalid, got ", invalidPassword)
		  } 
	  })
	t.Run("CreateValidPerson", func(t *testing.T) { 

		// It is not easy to pass a structure to an HTTP call!	
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(neilsBohr)

		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}

		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		exists := jsonParsed.ExistsP("data")
		if !exists  {
			t.Error("Expected Person data, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.Name")
		if !exists  {
			t.Error("Expected Person data.Name, dit not get ", "it")
		}
		returnedPersonName, ok := jsonParsed.Path("data.Name").Data().(string) 
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		} 
		if returnedPersonName != neilsBohr.Name {
			t.Errorf("Expected new Person Name to be %s, got %s", neilsBohr.Name, returnedPersonName)
		}
		returnedPersonEmail, ok := jsonParsed.Path("data.Email").Data().(string) 
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		} 
		if  returnedPersonEmail !=  strings.ToLower(neilsBohr.Email) {
			t.Errorf("Expected new Person Email to be %s, got %s", strings.ToLower(neilsBohr.Email), returnedPersonEmail)
		}
		returnedPersonPassword, ok := jsonParsed.Path("data.Password").Data().(string) 
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		} 
		if returnedPersonPassword != neilsBohr.Password {
			t.Errorf("Expected new Person Password to be %s, got%s", neilsBohr.Password, returnedPersonPassword)
		}
		exists = jsonParsed.ExistsP("data.CreatedAt")
		if !exists  {
			t.Error("Expected Person CreatedAt, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.UpdatedAt")
		if !exists  {
			t.Error("Expected Person UpdatedAt, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.DeletedAt")
		if !exists  {
			t.Error("Expected Person DeletedAt, dit not get ", "it")
		}
	})

	t.Run("GetPersonValidId", func(t *testing.T) { 
	})

	// tear down logic
}

// Sets a person table for testing by deleting all records in it, and adding two new records to it
// Input: 
// Output: 
//
func setupPersonTable() {
	// Clean up the table
	result := configs.DB.Exec("DELETE FROM people")
	if result.Error != nil {
		panic("unable to empty the person table")
	}

	// Insert two records
	result = configs.DB.Create(&albertEinstein)
	if result.Error != nil {
		panic("unable to add record to person table")
	}

	result = configs.DB.Create(&francisBacon)
	if result.Error != nil {
		panic("unable to add record to person table")
	}
}

// Serializes a structure
// Input: models.Person
// Output: *bytes.Buffer
// 
func serializePerson(person models.Person) *bytes.Buffer {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(person)
	return reqBodyBytes
}

// Creates a copy of Person structure, changes the attribute to be attribute value, and serializes the new structure
// Input: models.Person, string, string
// Output: *bytes.Buffer
// 
func getPersonWithInvalidAttributeSerialized(original models.Person, attribute string, attributeValue string) *bytes.Buffer {
	invalidPerson := original
	switch attribute {
		case "email":
			invalidPerson.Email = attributeValue
		case "name":
			invalidPerson.Name = attributeValue
		case "password":
			invalidPerson.Password = attributeValue
		default:	
	}
	return serializePerson(invalidPerson)
}