package assignment

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	validation "github.com/RodrigoMattosoSilveira/gradeit/controllers/validation"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	services "github.com/RodrigoMattosoSilveira/gradeit/services/assignment"
)

type controller struct {
	services services.AssignmentSvcInt
}

// NewAssignment - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func NewAssignment(s services.AssignmentSvcInt) controller {
	return controller{services: s}
}

func (c controller) Create(ctx *gin.Context) {
	var body models.AssignmentCreate
	errors := make([]string, 0)

	// bind request values to body and to assignment
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}
	assignment := models.Assignment{ID: uint64(body.PersonId), Description: body.Description, Due: body.Due}

	// ensure assignee, person_id, is the database
	if !PersonInDB(body.PersonId) {
		errors = append(errors, fmt.Sprintf("Assignment Create %d: person_id not in db", body.PersonId))
	}

	// ensure that the assignment's due date is beyond NOW
	if !DueNotOld(body.Due) {
		errors = append(errors, fmt.Sprintf("Assignment Create %d: invalid due date", body.PersonId))
	}

	// Insert new record in the dadtabasem return otherwise
	if len(errors) > 0 {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}
	c.services.Create(ctx, assignment)
}


func (c controller) GetAll(ctx *gin.Context) {
	c.services.GetAll(ctx)
}

func (c controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.services.GetByID(ctx, idInt)
}

func (c controller) Update(ctx *gin.Context) {
	var body models.AssignmentCreate
	errors := make([]string, 0)

	// bind request values to body and to assignment
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}
	assignment := models.Assignment{ID: uint64(body.PersonId), Description: body.Description, Due: body.Due}

	// Retrieve the ID of the assignment to be updated
	idParm, err := validation.ParseIdParm(ctx)
	if !err  {
		errors = append(errors, "Assignment Update, unable to parse  assignment ID")
	}

	//  ensure assignment id is valid
	id, valid := ValidIdParm(idParm)
	if !valid {
		errors = append(errors, fmt.Sprintf("Assignment Update %d: invalid assignment id", id))
	}

	//  ensure assignment id is in the db
	if !AssignmentInDB(uint64(id)) {
		errors = append(errors, fmt.Sprintf("Assignment Update %d: assignment not in db", id))
	}

	if !body.Due.IsZero() && !DueNotOld(assignment.Due) {
		errors = append(errors, fmt.Sprintf("Assignment Update %d: due date is not beyond NOW", id))
	}

	if len(errors) > 0 {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}
	c.services.Update(ctx, assignment)
}

func (c controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.services.Delete(ctx, idInt)
}