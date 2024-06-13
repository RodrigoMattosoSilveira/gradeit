package person

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	validation "github.com/RodrigoMattosoSilveira/gradeit/controllers/validation"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	"github.com/RodrigoMattosoSilveira/gradeit/services/person"

	"github.com/gin-gonic/gin"
)

type controller struct {
	services person.PersonSvcInt
}

// NewPerson - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func NewPerson(s person.PersonSvcInt) controller {
	return controller{services: s}
}

func (c controller) Create(ctx *gin.Context) {
	var body models.PersonCreate
	valid := true
	var personValidation models.PersonValidation

	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&body); err != nil {
			if err := ctx.ShouldBind(&body); err != nil {
				slog.Error("PersonCreate: Unable to determine the request content")
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "PersonCreate: Unable to determine the request content-type"})
				return
			}
		}
	}
	person := models.Person{Name: body.Name, Email: strings.ToLower(body.Email), Password: body.Password}

	if !ValidEmail(person.Email) {
		valid = false
		personValidation.InvalidEmail = true
		slog.Error(fmt.Sprintf("PersonCreate %d: invalid email = %s", person.ID, person.Email))
	}

	if valid && !UniqueEmail(person.Email) {
		valid = false
		personValidation.EmailExists = true
		slog.Error(fmt.Sprintf("PersonCreate %d: email already exists = %s", person.ID, person.Password))
	}

	if !ValidPassword(person.Password) {
		valid = false
		personValidation.InvalidPassword = true
		slog.Error(fmt.Sprintf("PersonCreate %d: invalid password = %s", person.ID, person.Password))
	}

	if !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{`error`: personValidation})
		return
	}
	c.services.Create(ctx, person)
}

func (c controller) GetAll(ctx *gin.Context) {
	c.services.GetAll(ctx)
}

func (c controller) GetByID(ctx *gin.Context) {
	valid := true
	var personValidation models.PersonValidation

	idParm, err := validation.ParseIdParm(ctx)
	if !err {
		valid = false
		personValidation.ParmIdInexistent = true
	}

	id, valid := ValidIdParm(idParm)
	if !valid {
		valid = false
		personValidation.InvalidParmId = true
	}

	if !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{`error`: personValidation})
		return
	}
	c.services.GetByID(ctx, id)
}

func (c controller) Update(ctx *gin.Context) {
	var body models.PersonUpdate
	valid := true
	var personValidation models.PersonValidation

	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&body); err != nil {
			if err := ctx.ShouldBind(&body); err != nil {
				slog.Error("PersonCreate: Unable to determine the request content")
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "PersonCreate: Unable to determine the request content-type"})
				return
			}
		}
	}
	personUpateData := models.Person{Name: body.Name, Email: strings.ToLower(body.Email), Password: body.Password}

	// Update approach
	// - Ensure at least one attribute is submtied for update
	// - Ensure that the person ID is valid, and there is a person with the ID in the database
	// - Validate the remaining attributes

	// At least one attribute must be submitted for update
	if personUpateData.Name == "" && personUpateData.Email == "" && personUpateData.Password == "" {
		personValidation.NoUpdateAttributes = true
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{`error`: personValidation})
		return
	}

	// Ensure that the person ID is valid, and there is a person with the ID in the database
	idParm, err := validation.ParseIdParm(ctx)
	if !err {
		valid = false
		personValidation.ParmIdInexistent = true
	}

	id, validParmId := ValidIdParm(idParm)
	if !validParmId {
		valid = false
		personValidation.InvalidParmId = true
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{`error`: personValidation})
		return
	} else {
		personUpateData.ID = uint64(id)
	}

	if !PersonInDB(personUpateData.ID) {
		valid = false
		personValidation.PersonNotInDB = true
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{`error`: personValidation})
		return
	}

	// Validate the remaining attributes
	if personUpateData.Email != "" && !ValidEmail(personUpateData.Email) {
		valid = false
		personValidation.InvalidEmail = true
		slog.Error(fmt.Sprintf("PersonCreate %d: invalid email = %s", personUpateData.ID, personUpateData.Email))
	}

	if valid && personUpateData.Email != "" && !UniqueEmail(personUpateData.Email) {
		valid = false
		personValidation.EmailExists = true
		slog.Error(fmt.Sprintf("PersonCreate %d: email already exists = %s", personUpateData.ID, personUpateData.Password))
	}

	if personUpateData.Password != "" && !ValidPassword(personUpateData.Password) {
		valid = false
		personValidation.InvalidPassword = true
		slog.Error(fmt.Sprintf("PersonCreate %d: invalid password = %s", personUpateData.ID, personUpateData.Password))
	}

	if !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{`error`: personValidation})
		return
	}
	c.services.Update(ctx, personUpateData)
}

func (c controller) Delete(ctx *gin.Context) {
	valid := true
	errorType := http.StatusOK
	var personValidation models.PersonValidation

	idParm, err := validation.ParseIdParm(ctx)
	if !err {
		valid = false
		personValidation.ParmIdInexistent = true
		errorType = http.StatusUnprocessableEntity
	}

	id, valid := ValidIdParm(idParm)
	if !valid {
		valid = false
		personValidation.InvalidParmId = true
		errorType = http.StatusUnprocessableEntity
	}

	if valid && !PersonInDB(id) {
		valid = false
		personValidation.PersonNotInDB = true
		errorType = http.StatusNotFound
	}

	if !valid {
		ctx.JSON(errorType, gin.H{`error`: personValidation})
		return
	}
	c.services.Delete(ctx, id)
}
