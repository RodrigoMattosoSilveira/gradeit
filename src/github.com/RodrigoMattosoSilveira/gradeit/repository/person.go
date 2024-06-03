package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/models"
	"github.com/RodrigoMattosoSilveira/gradeit/interfaces"
)

type repository struct{}

// NewPerson is a factory function for store layer that returns a interface type, UserInt
func NewPerson() interfaces.PersonCrudInt {
	return repository{}
}

// A RepositoryInt interface method
//
// Inserts a record in the user table
func (repo repository) Create(ctx *gin.Context, person models.Person) {
	ctx.JSON(200, gin.H{"data": "person created"})
}

func (repo repository) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": "people retrieved"})
}

func (repo repository) GetByID(ctx *gin.Context, id uint64) {
	ctx.JSON(200, gin.H{"data": "person retrieved"})
}

func (repo repository) Update(ctx *gin.Context, id uint64, person models.Person) {
	ctx.JSON(200, gin.H{"data": "person updated"})
}
func (repo repository) Delete(ctx *gin.Context, id uint64) {
	ctx.JSON(200, gin.H{"data": "person deleted"})
}