package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type AssignmentRepoInt interface {
	Create(ctx *gin.Context, assignment models.Assignment)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id uint64)
	Update(ctx *gin.Context, assignment models.Assignment)
	Delete(ictx *gin.Context, id uint64)
}