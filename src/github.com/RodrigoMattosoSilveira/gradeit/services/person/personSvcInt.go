package person

import (
	"github.com/gin-gonic/gin"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type PersonSvcInt interface {
	Create(ctx *gin.Context, entity models.Person)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id uint64)
	Update(ctx *gin.Context,entity models.Person)
	Delete(ctx *gin.Context, id uint64)
}