package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type repository struct{}

// NewAssignment is a factory function for store layer that returns a interface type
func NewAssignment() AssignmentRepoInt {
	return repository{}
}

// cURL validation command
// curl -X POST --json '{"person_id": 1, "description": "Find 4th law", "due": "2025-10-31T09:00:00.594Z"}' localhost:${HTTP_PORT}/assignment
// 
func (repo repository) Create(ctx *gin.Context, assignment models.Assignment) {
	result := configs.DB.Create(&assignment)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": assignment})
}

// cURL validation command
// curl -X GET localhost:${HTTP_PORT}/assignment
//
func (repo repository) GetAll(ctx *gin.Context) {
	var assignments []models.Assignment

	result := configs.DB.Find(&assignments)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
	return
	}

	ctx.JSON(200, gin.H{"data": assignments})
}

// cURL validation command
// curl -X GET localhost:${HTTP_PORT}/assignment/1
//
func (repo repository) GetByID(ctx *gin.Context, id uint64) {
	var assignment models.Assignment

	result := configs.DB.First(& assignment, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data":  assignment})
}


// cURL validation command
// curl -X PUT --json '{"person_id": 1, "description": "Find 4th law", "due": "2025-10-31T09:00:00.594Z"}' localhost:${HTTP_PORT}/assignment/1
// 
func (repo repository) Update(ctx *gin.Context, entity models.Assignment) {
	configs.DB.Model(&entity).Updates(models.Assignment{PersonId: entity.PersonId, Description: entity.Description, Due: entity.Due})

	ctx.JSON(200, gin.H{"data": entity})
}

// cURL validation command
// curl -X DELETE localhost:${HTTP_PORT}/assignment/1
//
func (repo repository) Delete(ctx *gin.Context, id uint64) {
	configs.DB.Delete(&models.Assignment{}, id)

	ctx.JSON(200, gin.H{"data": fmt.Sprintf("deleted assignment.id=%d", id)})
}