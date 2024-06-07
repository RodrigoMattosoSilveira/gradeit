package main

import (
	// "net/http"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/stretchr/testify/assert"
	cfg "github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

//  Inspired by
// https://blog.marcnuri.com/go-testing-gin-gonic-with-httptest
// 
func TestPersonCreateOK(t *testing.T) {
	router := GetRouter()
	RoutesPerson(router)
	cfg.Config()

	person := models.PersonCreate {
		Name: "Thomas Edison",
		Email: "edison@mail.com",
		Password: "edison123",
	}

	recorder := httptest.NewRecorder()
	personString := `{"name": "Thomas Edison", "email": "edison@mail.com", "password": "edison124"}`
	router.ServeHTTP(recorder, httptest.NewRequest("POST", "/person", strings.NewReader(personString)))

	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
		  t.Error("Expected 200, got ", recorder.Code)
		}
	})

	var personDataMap map[string]map[string]interface{}
	err := json.Unmarshal([]byte(recorder.Body.Bytes()), &personDataMap)
	if err != nil {
		panic(err)
	}

	t.Run("Returns person with id = 1", func(t *testing.T) {
		if _, isMapContainsKey := personDataMap["data"]["CreatedAt"]; isMapContainsKey {
		//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected CreatedAt attribute, got no CreatedAt attribute"))
		}
		if _, isMapContainsKey := personDataMap["data"]["UpdatedAt"]; isMapContainsKey {
			//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected UpdatedAt attribute, got no CreatedAt attribute"))
		}
		if _, isMapContainsKey := personDataMap["data"]["DeletedAt"]; isMapContainsKey {
			//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected DeletedAt attribute, got no DeletedAt attribute"))
		}
		if _, isMapContainsKey := personDataMap["data"]["Name"]; isMapContainsKey {
			//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected Name attribute, got no Name attribute"))
		}
		if personDataMap["data"]["Name"] != person.Name {
			t.Error(fmt.Errorf("Expected Person Name: %s, got %s", person.Name, personDataMap["Name"]))
		}
		if personDataMap["data"]["Email"] != person.Email {
			t.Error(fmt.Errorf("Expected Persob Email: %s, got %s", person.Email, personDataMap["Email"]))
		}
		if personDataMap["data"]["Password"] != person.Password {
			t.Error(fmt.Errorf("Expected Person Password: %s, got %s", person.Password, personDataMap["Password"]))
		}
	})
}

func TestPersonGetIdOK(t *testing.T) {
	router := GetRouter()
	RoutesPerson(router)
	cfg.Config()

	person := models.Person {
		Name: "Albert Einstein",
		Email: "einstein@gmail.com",
		Password: "einstein124",
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/person/1", nil))

	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
		  t.Error("Expected 200, got ", recorder.Code)
		}
	})

	var personDataMap map[string]map[string]interface{}
	err := json.Unmarshal([]byte(recorder.Body.Bytes()), &personDataMap)
	if err != nil {
		panic(err)
	}

	t.Run("Returns person with id = 1", func(t *testing.T) {
		if _, isMapContainsKey := personDataMap["data"]["CreatedAt"]; isMapContainsKey {
		//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected CreatedAt attribute, got no CreatedAt attribute"))
		}
		if _, isMapContainsKey := personDataMap["data"]["UpdatedAt"]; isMapContainsKey {
			//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected UpdatedAt attribute, got no CreatedAt attribute"))
		}
		if _, isMapContainsKey := personDataMap["data"]["DeletedAt"]; isMapContainsKey {
			//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected DeletedAt attribute, got no DeletedAt attribute"))
		}
		if _, isMapContainsKey := personDataMap["data"]["Name"]; isMapContainsKey {
			//key exist
		} else {
			//key does not exist
			t.Error(fmt.Errorf("Expected Name attribute, got no Name attribute"))
		}
		if personDataMap["data"]["Name"] != person.Name {
			t.Error(fmt.Errorf("Expected Person Name: %s, got %s", person.Name, personDataMap["Name"]))
		}
		if personDataMap["data"]["Email"] != person.Email {
			t.Error(fmt.Errorf("Expected Persob Email: %s, got %s", person.Email, personDataMap["Email"]))
		}
		if personDataMap["data"]["Password"] != person.Password {
			t.Error(fmt.Errorf("Expected Person Password: %s, got %s", person.Password, personDataMap["Password"]))
		}
	})
}

func TestPersonGetIdNotOK(t *testing.T) {
	router := GetRouter()
	RoutesPerson(router)
	cfg.Config()

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/person/10000", nil))

	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 422 {
		  t.Error("Expected 200, got ", recorder.Code)
		}
	})
}

// Convert a structure to a string
// Input: interface{}
// Output: (string, nil) if OK,  ("", error) otherwise 
// 
func structToString (sourceStruct interface{}) (string, error) {
	sourceJson, error := json.Marshal(sourceStruct)
	if error != nil {
		return "", error
	} 
	return string(sourceJson), nil
}