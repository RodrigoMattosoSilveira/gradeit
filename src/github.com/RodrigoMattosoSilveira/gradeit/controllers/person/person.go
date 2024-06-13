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
	person := models.Person{Name: body.Name, Email: strings.ToLower(body.Email), Password: body.Password}

	idParm, err := validation.ParseIdParm(ctx)
	if !err {
		valid = false
		personValidation.ParmIdInexistent = true
	}

	id, valid := ValidIdParm(idParm)
	if !valid {
		valid = false
		personValidation.InvalidParmId = true
	} else {
		person.ID = id
	}

	if !PersonInDB(uint64(person.ID)) {
		valid = false
		personValidation.PersonNotInDB = true
	}

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
	c.services.Update(ctx, person)
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
