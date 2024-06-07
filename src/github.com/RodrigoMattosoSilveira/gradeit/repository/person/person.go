package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type repository struct{}

// NewPerson is a factory function for store layer that returns a interface type, UserInt
func NewPerson() PersonRepoInt {
	return repository{}
}

// cURL validation command, $ export HTTP_PORT=<<port service is listening on>>
// curl -X POST --json '{"name": "Albert Einstein", "email": "einstein@mail.com", "password": "einstein124"}' localhost:${HTTP_PORT}/person
// 
func (repo repository) Create(ctx *gin.Context, person models.Person) {
	result := configs.DB.Create(&person)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

// cURL validation command, $ export HTTP_PORT=<<port service is listening on>>
// curl -X GET localhost:${HTTP_PORT}/person
// 
func (repo repository) GetAll(ctx *gin.Context) {
	var people []models.Person

	result := configs.DB.Find(&people)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": people})
}

// cURL validation command, $ export HTTP_PORT=<<port service is listening on>>
// curl -X GET localhost:${HTTP_PORT}/person/1
// 
func (repo repository) GetByID(ctx *gin.Context, id uint64) {
	var person models.Person

	result := configs.DB.First(&person, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

// cURL validation command, $ export HTTP_PORT=<<port service is listening on>>
// curl -X PUT  --json '"email": "einstein.new.email@mail.com"}' localhost:${HTTP_PORT}/person/1
// 
func (repo repository) Update(ctx *gin.Context, person models.Person) {
	configs.DB.Model(&person).Updates(models.Person{Name: person.Name, Email: person.Email, Password: person.Password})

	ctx.JSON(200, gin.H{"data": person})
}

// cURL validation command, $ export HTTP_PORT=<<port service is listening on>>
// curl -X DELETE localhost:${HTTP_PORT}/person/1
// 
func (repo repository) Delete(ctx *gin.Context, id uint64) {
	configs.DB.Delete(&models.Person{}, id)

	ctx.JSON(200, gin.H{"data": fmt.Sprintf("deleted person.id=%d", id)})
}
