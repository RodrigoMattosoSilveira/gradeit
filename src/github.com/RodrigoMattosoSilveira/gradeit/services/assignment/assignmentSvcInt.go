package assignment

import (
	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type AssignmentSvcInt interface {
	Create(ctx *gin.Context, entity models.Assignment)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id uint64)
	Update(ctx *gin.Context,entity models.Assignment)
	Delete(ctx *gin.Context, id uint64)
}