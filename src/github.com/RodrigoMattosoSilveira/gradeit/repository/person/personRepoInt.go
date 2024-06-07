package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type PersonRepoInt interface {
	Create(ctx *gin.Context, person models.Person)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id uint64)
	Update(ctx *gin.Context, person models.Person)
	Delete(ictx *gin.Context, id uint64)
}