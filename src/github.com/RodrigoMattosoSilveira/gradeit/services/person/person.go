package person

import (
	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/gradeit/models"
)

type service struct {
	repository PersonSvcInt
}

// NewPerson - is a factory function to inject store in service.
func NewPerson(s PersonSvcInt) PersonSvcInt {
	return service{repository: s}
}

func (s service) Create(ctx *gin.Context, person models.Person) {
	s.repository.Create(ctx, person)
}

func (s service) GetAll(ctx *gin.Context) {
	s.repository.GetAll(ctx)
}

func (s service) GetByID(ctx *gin.Context, id uint64) {
	s.repository.GetByID(ctx, id)
}

func (s service) Update(ctx *gin.Context, person models.Person) {
	s.repository.Update(ctx, person)
}

func (s service) Delete(ctx *gin.Context, id uint64) {
	s.repository.Delete(ctx, id)
}
