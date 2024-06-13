package assignment

import (
	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type service struct {
	repository AssignmentSvcInt
}

// NewAssignment - is a factory function to inject store in service.
func NewAssignment(a AssignmentSvcInt) AssignmentSvcInt {
	return service{repository: a}
}

func (s service) Create(ctx *gin.Context, Assignment models.Assignment) {
	s.repository.Create(ctx, Assignment)
}

func (s service) GetAll(ctx *gin.Context) {
	s.repository.GetAll(ctx)
}

func (s service) GetByID(ctx *gin.Context, id uint64) {
	s.repository.GetByID(ctx, id)
}

func (s service) Update(ctx *gin.Context, Assignment models.Assignment) {
	s.repository.Update(ctx, Assignment)
}

func (s service) Delete(ctx *gin.Context, id uint64) {
	s.repository.Delete(ctx, id)
}