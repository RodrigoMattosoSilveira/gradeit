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

var antoineLavoisier = models.Person {
	Name: "Antoine Lavoisier",
	Email:  strings.ToLower("Antoine.Lavoisier@gmail.com"),
	Password: "Antoine.Lavoisier123",
}

var neilsBohr = models.Person {
	Name: "Neils Bohr",
	Email:  strings.ToLower("Neils.Bohrn@gmail.com"),
	Password: "Neils.Bohr123",
}

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
		neilsBohrWork := neilsBohr
		neilsBohrWork.Email =  strings.ToLower("Neils.Bohrn@gm")
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(neilsBohrWork)

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
		neilsBohrWork := neilsBohr
		neilsBohrWork.Email =  strings.ToLower("Albert.Einstein@gmail.com")
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(neilsBohrWork)

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
		neilsBohrWork := neilsBohr
		neilsBohrWork.Password = "Nei"
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(neilsBohrWork)

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

// //  Inspired by
// // https://blog.marcnuri.com/go-testing-gin-gonic-with-httptest
// //
// func TestPersonCreateOK(t *testing.T) {
// 	configs.Config()
// 	router := configs.GetRouter()
// 	routes.RoutesPerson(router)


// 	// Delete all records with the email of the record about to be inserted
// 	configs.DB.Exec("DELETE FROM people WHERE email like 'Einstein123abc'")

// 	person := models.PersonCreate {
// 		Name: "Neils Bohr",
// 		Email: "Bohr@mail.com",
// 		Password: "Bohr123abc",
// 	}

// 	recorder := httptest.NewRecorder()
// 	personString := `{"Name": " Albert Einstein", "Email": "Einstein@mail.com", "Password": "Einstein123abc"}`
// 	req, _ := http.NewRequest("POST", "/person", strings.NewReader(personString))
// 	req.Header.Set("Content-Type", "application/json")

//  	router.ServeHTTP(recorder, req)


// 	t.Run("Returns 200 status code", func(t *testing.T) {
// 		if recorder.Code != 200 {
// 		  t.Error("Expected 200, got ", recorder.Code)
// 		}
// 	})

// 	var personDataMap map[string]map[string]interface{}
// 	err := json.Unmarshal([]byte(recorder.Body.Bytes()), &personDataMap)
// 	if err != nil {
// 		panic(err)
// 	}

// 	t.Run("Returns person with id = 1", func(t *testing.T) {
// 		if _, isMapContainsKey := personDataMap["data"]["CreatedAt"]; isMapContainsKey {
// 		//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected CreatedAt attribute, got no CreatedAt attribute"))
// 		}
// 		if _, isMapContainsKey := personDataMap["data"]["UpdatedAt"]; isMapContainsKey {
// 			//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected UpdatedAt attribute, got no CreatedAt attribute"))
// 		}
// 		if _, isMapContainsKey := personDataMap["data"]["DeletedAt"]; isMapContainsKey {
// 			//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected DeletedAt attribute, got no DeletedAt attribute"))
// 		}
// 		if _, isMapContainsKey := personDataMap["data"]["Name"]; isMapContainsKey {
// 			//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected Name attribute, got no Name attribute"))
// 		}
// 		if personDataMap["data"]["Name"] != person.Name {
// 			t.Error(fmt.Errorf("Expected Person Name: %s, got %s", person.Name, personDataMap["Name"]))
// 		}
// 		if personDataMap["data"]["Email"] != person.Email {
// 			t.Error(fmt.Errorf("Expected Persob Email: %s, got %s", person.Email, personDataMap["Email"]))
// 		}
// 		if personDataMap["data"]["Password"] != person.Password {
// 			t.Error(fmt.Errorf("Expected Person Password: %s, got %s", person.Password, personDataMap["Password"]))
// 		}
// 	})
// }

// func TestPersonCreateEmailExists(t *testing.T) {
// 	configs.Config()
// 	router := configs.GetRouter()
// 	routes.RoutesPerson(router)


// 	recorder := httptest.NewRecorder()
// 	personString := `{"Name": "Richard Feyman", "Email": "Bohr@mail.com", "Password": "Feyman123abc"}`
// 	req, _ := http.NewRequest("POST", "/person", strings.NewReader(personString))
// 	req.Header.Set("Content-Type", "application/json")

//  	router.ServeHTTP(recorder, req)


// 	t.Run("Returns 200 status code", func(t *testing.T) {
// 		if recorder.Code != 422 {
// 		  t.Error("Expected 422, got ", recorder.Code)
// 		}
// 	})
// }

// func TestPersonGetIdOK(t *testing.T) {
// 	configs.Config()
// 	router := configs.GetRouter()
// 	routes.RoutesPerson(router)

// 	person := models.Person {
// 		Name: "Albert Einstein",
// 		Email: "einstein@gmail.com",
// 		Password: "einstein124",
// 	}

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/person/1", nil))

// 	t.Run("Returns 200 status code", func(t *testing.T) {
// 		if recorder.Code != 200 {
// 		  t.Error("Expected 200, got ", recorder.Code)
// 		}
// 	})

// 	var personDataMap map[string]map[string]interface{}
// 	err := json.Unmarshal([]byte(recorder.Body.Bytes()), &personDataMap)
// 	if err != nil {
// 		panic(err)
// 	}

// 	t.Run("Returns person with id = 1", func(t *testing.T) {
// 		if _, isMapContainsKey := personDataMap["data"]["CreatedAt"]; isMapContainsKey {
// 		//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected CreatedAt attribute, got no CreatedAt attribute"))
// 		}
// 		if _, isMapContainsKey := personDataMap["data"]["UpdatedAt"]; isMapContainsKey {
// 			//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected UpdatedAt attribute, got no CreatedAt attribute"))
// 		}
// 		if _, isMapContainsKey := personDataMap["data"]["DeletedAt"]; isMapContainsKey {
// 			//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected DeletedAt attribute, got no DeletedAt attribute"))
// 		}
// 		if _, isMapContainsKey := personDataMap["data"]["Name"]; isMapContainsKey {
// 			//key exist
// 		} else {
// 			//key does not exist
// 			t.Error(fmt.Errorf("Expected Name attribute, got no Name attribute"))
// 		}
// 		if personDataMap["data"]["Name"] != person.Name {
// 			t.Error(fmt.Errorf("Expected Person Name: %s, got %s", person.Name, personDataMap["Name"]))
// 		}
// 		if personDataMap["data"]["Email"] != person.Email {
// 			t.Error(fmt.Errorf("Expected Persob Email: %s, got %s", person.Email, personDataMap["Email"]))
// 		}
// 		if personDataMap["data"]["Password"] != person.Password {
// 			t.Error(fmt.Errorf("Expected Person Password: %s, got %s", person.Password, personDataMap["Password"]))
// 		}
// 	})
// }

// func TestPersonGetIdNotOK(t *testing.T) {
// 	configs.Config()
// 	router := configs.GetRouter()
// 	routes.RoutesPerson(router)

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/person/10000", nil))

// 	t.Run("Returns 200 status code", func(t *testing.T) {
// 		if recorder.Code != 422 {
// 		  t.Error("Expected 200, got ", recorder.Code)
// 		}
// 	})
// }