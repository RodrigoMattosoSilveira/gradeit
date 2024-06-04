package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/interfaces"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
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
	result := configs.DB.Create(&person)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

func (repo repository) GetAll(ctx *gin.Context) {
	var people []models.Person

	result := configs.DB.Find(&people)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": people})
}

func (repo repository) GetByID(ctx *gin.Context, id uint64) {
	var person models.Person

	result := configs.DB.First(&person, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

func (repo repository) Update(ctx *gin.Context, person models.Person) {
	configs.DB.Model(&person).Updates(models.Person{Name: person.Name, Email: person.Email, Password: person.Password})

	ctx.JSON(200, gin.H{"data": person})
}
func (repo repository) Delete(ctx *gin.Context, id uint64) {
	configs.DB.Delete(&models.Person{}, id)

	ctx.JSON(200, gin.H{"data": fmt.Sprintf("deleted person.id=%d", id)})
}