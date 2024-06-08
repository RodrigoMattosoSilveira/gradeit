package person

import (
	"fmt"

	"github.com/RodrigoMattosoSilveira/gradeit/services/person"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	validation "github.com/RodrigoMattosoSilveira/gradeit/controllers/validation"

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
	errors := make([]string, 0)
	var body models.PersonCreate

	contentType := ctx.Request.Header.Get("Content-Type")
	if (contentType == "application/json") {
		if err := ctx.ShouldBindJSON(&body); err != nil {
			if err := ctx.ShouldBind(&body); err != nil {
				errors = append(errors, err.Error())
				ctx.JSON(422, gin.H{"error": errors})
				return
			}
		}
	}
	person := models.Person{Name: body.Name, Email: body.Email, Password: body.Password}

	if !ValidEmail(person.Email) {
		errors = append(errors, fmt.Sprintf("PersonCreate %d: invalid email = %s", person.ID, person.Email))
	}

	if !UniqueEmail(person.Email) {
		errors = append(errors, fmt.Sprintf("PersonCreate %d: email already exists = %s", person.ID, person.Password))
	}

	if !ValidPassword(person.Password) {
		errors = append(errors, fmt.Sprintf("PersonCreate %d: invalid password = %s", person.ID, person.Password))
	}
	if len(errors) > 0 {
		ctx.JSON(422, gin.H{"error": errors})
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
