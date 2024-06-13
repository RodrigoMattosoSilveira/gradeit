package main

import (
	// "net/http"
	// "encoding/json"
	// "fmt"
	"bytes"
	"encoding/json"
	"fmt"
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

// Relying on https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses
// 
var albertEinstein = models.Person{
	Name:     "Albert Einstein",
	Email:    strings.ToLower("Albert.Einstein@gmail.com"),
	Password: "Albert.Einstein123",
}

var francisBacon = models.Person{
	Name:     "Francis Bacon",
	Email:    strings.ToLower("Francis.Bacon@gmail.com"),
	Password: "Francis.Bacon123",
}

// var antoineLavoisier = models.PersonCreate {
// 	Name: "Antoine Lavoisier",
// 	Email:  strings.ToLower("Antoine.Lavoisier@gmail.com"),
// 	Password: "Antoine.Lavoisier123",
// }

var neilsBohr = models.Person{
	Name:     "Neils Bohr",
	Email:    strings.ToLower("Neils.Bohrn@gmail.com"),
	Password: "Neils.Bohr123",
}

//	Inspired by
//
// // https://blog.marcnuri.com/go-testing-gin-gonic-with-httptest
func TestPerson(t *testing.T) {
	// setup logic
	configs.Config()
	router := configs.GetRouter()
	routes.RoutesPerson(router)
	setupPersonTable()
	t.Run("CreateInvalidEmail", func(t *testing.T) {

		// Set up a person structure with an invalid email
		invalidPerson := neilsBohr
		invalidPerson.Email = "bad@email"

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(invalidPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
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

		// Set up a person structure with an existing email
		invalidPerson := neilsBohr
		invalidPerson.Email = "Albert.Einstein@gmail.com"

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(invalidPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		// Validate the results
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

		// Set up a person structure with an invalid password
		invalidPerson := neilsBohr
		invalidPerson.Password = "Nei"

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(invalidPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		// Validate the results
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

		// Serialize the valid person structure
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(neilsBohr)

		// Set up the HTTP call
		req, _ := http.NewRequest("POST", "/person", reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected %d got %d", http.StatusCreated, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		exists := jsonParsed.ExistsP("data")
		if !exists {
			t.Error("Expected Person data, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.Name")
		if !exists {
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
		if returnedPersonEmail != strings.ToLower(neilsBohr.Email) {
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
		if !exists {
			t.Error("Expected Person CreatedAt, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.UpdatedAt")
		if !exists {
			t.Error("Expected Person UpdatedAt, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.DeletedAt")
		if !exists {
			t.Error("Expected Person DeletedAt, dit not get ", "it")
		}
	})
	t.Run("GetInvalidParmId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with an invalid ID paramater
		req, _ := http.NewRequest("GET", fmt.Sprintf("/person/%s", "1A"), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		invalidPassword, ok := jsonParsed.Path("error.invalid_parm_id").Data().(bool)
		if !ok {
			panic("GetId Unable to parse the error")
		}
		if !invalidPassword {
			t.Error("Person GetId Expected Invalid Parameter Id error, got ", "something else")
		}
	})

	t.Run("GetValidId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with a valid ID
		req, _ := http.NewRequest("GET", fmt.Sprintf("/person/%d", albertEinstein.ID), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code !=  http.StatusOK {
			t.Errorf("Expected %d got %d", http.StatusOK, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		// We got the right name
		returnedPersonName, ok := jsonParsed.Path("data.Name").Data().(string)
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		}
		if returnedPersonName != albertEinstein.Name {
			t.Errorf("Expected new Person Name to be %s, got %s", albertEinstein.Name, returnedPersonName)
		}
		// We got the right email
		returnedPersonEmail, ok := jsonParsed.Path("data.Email").Data().(string)
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		}
		if returnedPersonEmail != strings.ToLower(albertEinstein.Email) {
			t.Errorf("Expected new Person Email to be %s, got %s", strings.ToLower(albertEinstein.Email), returnedPersonEmail)
		}
		// We got the right password
		returnedPersonPassword, ok := jsonParsed.Path("data.Password").Data().(string)
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		}
		if returnedPersonPassword != albertEinstein.Password {
			t.Errorf("Expected new Person Password to be %s, got%s", albertEinstein.Password, returnedPersonPassword)
		}
	})
	t.Run("GetIdNotInDb", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with an invalid ID paramater
		id := albertEinstein.ID - 1
		if albertEinstein.ID == 1 {
			id = 100
		} 
		req, _ := http.NewRequest("GET", fmt.Sprintf("/person/%d", id), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotFound {
			t.Errorf("Expected %d got %d", http.StatusNotFound, recorder.Code)
		}
	})

	t.Run("Get", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with an invalid ID paramater
		req, _ := http.NewRequest("GET", "/person", reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected %d got %d", http.StatusOK, recorder.Code)
		}

		// Validate the results
	})
	t.Run("UpdateNoParmId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(albertEinstein)

		// Set up the HTTP call with an invalid ID paramater
		req, _ := http.NewRequest("PUT", "/person", reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotFound {
			t.Errorf("Expected %d got %d", http.StatusNotFound, recorder.Code)
		}
	})
	t.Run("UpdateInvalidParmId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(albertEinstein)

		// Set up the HTTP call with an invalid ID paramater
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/person/%s", "1A"), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		invalidPassword, ok := jsonParsed.Path("error.invalid_parm_id").Data().(bool)
		if !ok {
			panic("Unable to parse the error")
		}
		if !invalidPassword {
			t.Error("Expected Invalid Parameter Id error, got ", "something else")
		}
	})
	t.Run("UpdateIdNotInDb", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(francisBacon)

		// Set up the HTTP call with an invalid ID paramater
		id := albertEinstein.ID - 1
		if albertEinstein.ID == 1 {
			id = 100
		} 
		req, _ := http.NewRequest("GET", fmt.Sprintf("/person/%d", id), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotFound {
			t.Errorf("Expected %d got %d", http.StatusNotFound, recorder.Code)
		}
	})
	t.Run("UpdateInvalidEmail", func(t *testing.T) {

		// Set up a person structure with an invalid email
		invalidPerson := albertEinstein
		invalidPerson.Email = "bad@email"

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(invalidPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/person/%d", albertEinstein.ID), reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Update Email Expected HTTP Status %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		invalidEmail, ok := jsonParsed.Path("error.invalid_email").Data().(bool)
		if !ok {
			panic("Update Email Unable to find Invalid Email path")
		}
		if !invalidEmail {
			t.Error("Update Email Expected Person email to be invalid, got ", invalidEmail)
		}
	})
	t.Run("UpdateName", func(t *testing.T) {

		// Set up a person structure with the new email
		updatedPerson := models.Person{
			Name: "Albert NoMiddleName Einstein@gmail",
		}

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(updatedPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/person/%d", albertEinstein.ID), reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
	
		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected %d got %d", http.StatusOK, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		exists := jsonParsed.ExistsP("data")
		if !exists {
			t.Error("Expected Person data, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.Name")
		if !exists {
			t.Error("Expected Person data.Name, dit not get ", "it")
		}
		returnedPersonName, ok := jsonParsed.Path("data.Name").Data().(string)
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		}
		if returnedPersonName != updatedPerson.Name {
			t.Errorf("Expected new Person Name to be %s, got %s", updatedPerson.Name, returnedPersonName)
		}
	})
	t.Run("UpdateEmail", func(t *testing.T) {

		// Set up a person structure with the new email
		updatedPerson := models.Person{
			Email: "albert.einstein.1879@gmail.com",
		}

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(updatedPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/person/%d", albertEinstein.ID), reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
	
		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected %d got %d", http.StatusOK, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		exists := jsonParsed.ExistsP("data")
		if !exists {
			t.Error("Expected Person data, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.Email")
		if !exists {
			t.Error("Expected Person data.Password, dit not get ", "it")
		}
		returnedPersonEmail, ok := jsonParsed.Path("data.Email").Data().(string)
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		}
		if returnedPersonEmail != updatedPerson.Email {
			t.Errorf("Expected new Person Email to be %s, got %s", updatedPerson.Email, returnedPersonEmail)
		}
	})
	t.Run("UpdatePassword", func(t *testing.T) {

		// Set up a person structure with the new email
		updatedPerson := models.Person{
			Password: "syvlMOXG5Pj**oH7",
		}

		// Serialize the person structure with the invalid email
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(updatedPerson)

		// Set up the HTTP call
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/person/%d", albertEinstein.ID), reqBodyBytes)
		recorder := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
	
		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected %d got %d", http.StatusOK, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		exists := jsonParsed.ExistsP("data")
		if !exists {
			t.Error("Expected Person data, dit not get ", "it")
		}
		exists = jsonParsed.ExistsP("data.Password")
		if !exists {
			t.Error("Expected Person data.Password, dit not get ", "it")
		}
		returnedPersonPassword, ok := jsonParsed.Path("data.Password").Data().(string)
		if !ok {
			t.Error("Expected valid error, got ", "bad error")
		}
		if returnedPersonPassword != updatedPerson.Password {
			t.Errorf("Expected new Person Name to be %s, got %s", updatedPerson.Password, returnedPersonPassword)
		}
	})
	t.Run("DeleteInexistentParmId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with an invalid ID paramater
		req, _ := http.NewRequest("DELETE", "/person", reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotFound {
			t.Errorf("Person DELETE Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		// Validate the results
		// Note that we do not have a route for this case, hence there is anything else to validate!
	})
	t.Run("DeleteInvalidParmId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with an invalid ID paramater
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/person/%s", "1A"), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnprocessableEntity {
			t.Errorf("Person DELETE Invalid Parameter Id Expected %d got %d", http.StatusUnprocessableEntity, recorder.Code)
		}

		// Validate the results
		jsonParsed, err := gabs.ParseJSON([]byte(recorder.Body.Bytes()))
		if err != nil {
			panic(err)
		}
		invalidPassword, ok := jsonParsed.Path("error.invalid_parm_id").Data().(bool)
		if !ok {
			panic("Person DELETE Invalid Parameter Id Unable to parse the error")
		}
		if !invalidPassword {
			t.Error("Person DELETE Expected Invalid Parameter Id error, got ", "something else")
		}
	})

	t.Run("DeleteIdNotInDb", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with an invalid ID paramater
		id := albertEinstein.ID - 1
		if albertEinstein.ID == 1 {
			id = 100
		} 
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/person/%d", id), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotFound {
			t.Errorf("Person DELETE Expected %d got %d", http.StatusNotFound, recorder.Code)
		}
	})

	t.Run("DeleteValidId", func(t *testing.T) {

		// Set up an empty Body
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(" ")

		// Set up the HTTP call with a valid ID
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/person/%d", albertEinstein.ID), reqBodyBytes)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Execute the HTTP call
		router.ServeHTTP(recorder, req)
		if recorder.Code !=  http.StatusOK {
			t.Errorf("Person DELETE Expected %d got %d", http.StatusOK, recorder.Code)
		}

		// Get the record and ensure its DeletedAt is not NULL
		var people []models.Person
		configs.DB.Unscoped().Where(fmt.Sprintf("id=%d", albertEinstein.ID)).Find(&people)
		if len(people) != 1 {
			t.Errorf("DELETE Unable to read deleted record id %d", albertEinstein.ID)
		}
		if people[0].ID != albertEinstein.ID {
			t.Errorf("DELETE Retrieved Deleted record id %d does not match actual deleted record id %d", people[0].ID, albertEinstein.ID)
		}
	})


	// tear down logic
}

// Sets a person table for testing by deleting all records in it, and adding two new records to it
// Input:
// Output:
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