package controllers

import (
	"net/http"
	"strconv"

	"github.com/RodrigoMattosoSilveira/gradeit/interfaces"
	"github.com/RodrigoMattosoSilveira/gradeit/models"

	"github.com/gin-gonic/gin"
)

type controller struct {
	services interfaces.PersonCrudInt
}

// NewPerson - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func NewPerson(s interfaces.PersonCrudInt) controller {
	return controller{services: s}
}

func (c controller) Create(ctx *gin.Context) {
	var body models.PersonCreate
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person := models.Person{Name: body.Name, Email: body.Email, Password: body.Password}
	validPerson, errors := ValidatePersonCreate(person)
	if !validPerson {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}
	c.services.Create(ctx, person)
}

func (c controller) GetAll(ctx *gin.Context) {
	c.services.GetAll(ctx)
}

func (c controller) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}
	c.services.GetByID(ctx, id)
}

func (c controller) Update(ctx *gin.Context) {
	var body models.PersonUpdate
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person := models.Person{Name: body.Name, Email: body.Email, Password: body.Password}
	c.services.Update(ctx, id, person)
}

func (c controller) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.services.Delete(ctx, id)
}
