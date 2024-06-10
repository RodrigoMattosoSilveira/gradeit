package person

import (
	"fmt"
	"log/slog"
	"strings"

	validation "github.com/RodrigoMattosoSilveira/gradeit/controllers/validation"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	"github.com/RodrigoMattosoSilveira/gradeit/services/person"

	"github.com/gin-gonic/gin"
	"github.com/Jeffail/gabs"
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
	valid := true
	var body models.PersonCreate
	jsonObj := gabs.New()


	contentType := ctx.Request.Header.Get("Content-Type")
	if (contentType == "application/json") {
		if err := ctx.ShouldBindJSON(&body); err != nil {
			if err := ctx.ShouldBind(&body); err != nil {
				slog.Error("PersonCreate: Unable to determine the request content")
				ctx.JSON(422, gin.H{"error": "PersonCreate: Unable to determine the request content-type"})
				return
			}
		}
	}
	person := models.Person{Name: body.Name, Email: strings.ToLower(body.Email), Password: body.Password}

	if !ValidEmail(person.Email) {
		valid = false
		jsonObj.Set(false, "Email", "Valid")
		slog.Error(fmt.Sprintf("PersonCreate %d: invalid email = %s", person.ID, person.Email))
	}

	if valid && !UniqueEmail(person.Email) {
		valid = false
		jsonObj.Set(false, "Email", "Unique")
		slog.Error(fmt.Sprintf("PersonCreate %d: email already exists = %s", person.ID, person.Password))
	}

	if !ValidPassword(person.Password) {
		valid = false
		jsonObj.Set(false, "Password", "Valid")
		slog.Error(fmt.Sprintf("PersonCreate %d: invalid password = %s", person.ID, person.Password))
	}

	if !valid {
		ctx.JSON(422, gin.H{"error": jsonObj.String()})
		return
	}
	c.services.Create(ctx, person)
}

func (c controller) GetAll(ctx *gin.Context) {
	c.services.GetAll(ctx)
}

func (c controller) GetByID(ctx *gin.Context) {
	errors := make([]string, 0)

	err, idParm := validation.ParseIdParm(ctx)
	if !err {
		errors = append(errors, "Person GetById, unable to parse id parameter")
	}

	id, valid := ValidIdParm(idParm)
	if !valid {
		errors = append(errors, fmt.Sprintf("GetByID %d: invalid id", id))
	}

	if !PersonInDB(uint64(id)) {
		errors = append(errors, fmt.Sprintf("Person GetByID %d: person not in db", id))
	}

	if len(errors) > 0 {
		ctx.JSON(422, gin.H{"error": errors})
		return
	}

	c.services.GetByID(ctx, id)
}

func (c controller) Update(ctx *gin.Context) {
	var body models.PersonUpdate
	errors := make([]string, 0)

	valid, idParm := validation.ParseIdParm(ctx)
	if !valid {
		errors = append(errors, "Person GetById,unable to parse id parameter")
	}

	if err := ctx.ShouldBind(&body); err != nil {
		errors = append(errors, err.Error())
	}
	person := models.Person{Name: body.Name, Email: body.Email, Password: body.Password}

	id, valid := ValidIdParm(idParm)
	if !valid {
		errors = append(errors, fmt.Sprintf("PersonUpdate %d: invalid id", person.ID))
	} else {
		person.ID = id
	}

	if !PersonInDB(uint64(person.ID)) {
		errors = append(errors, fmt.Sprintf("PersonUpdate %s: person not in db", idParm))
	}

	if person.Email != "" && !ValidEmail(person.Email) {
		errors = append(errors, fmt.Sprintf("PersonUpdate %d: invalid email = %s", person.ID, person.Email))
	}

	if person.Password != "" && ValidPassword(person.Password) {
		errors = append(errors, fmt.Sprintf("PersonUpdate %d: invalid password = %s", person.ID, person.Password))
	}
	if len(errors) > 0 {
		ctx.JSON(422, gin.H{"error": errors})
		return
	}

	c.services.Update(ctx, person)
}

func (c controller) Delete(ctx *gin.Context) {
	errors := make([]string, 0)

	valid, idParm := validation.ParseIdParm(ctx)
	if !valid {
		errors = append(errors, "Person Delete, unable to parse id parameter")
	}

	id, valid := ValidIdParm(idParm)
	if !valid {
		errors = append(errors, fmt.Sprintf("Delete %d: invalid id", id))
	}

	if !PersonInDB(uint64(id)) {
		errors = append(errors, fmt.Sprintf("Delete %s: person not in db", idParm))
	}

	if len(errors) > 0 {
		ctx.JSON(422, gin.H{"error": errors})
		return
	}

	c.services.Delete(ctx, id)
}
