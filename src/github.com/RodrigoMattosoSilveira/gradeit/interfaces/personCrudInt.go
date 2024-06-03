package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type PersonCrudInt interface {
	Create(ctx *gin.Context, person models.Person)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id uint64)
	Update(ctx *gin.Context, id uint64, person models.Person)
	Delete(ctx *gin.Context, id uint64)
}