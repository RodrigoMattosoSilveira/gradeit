package main

import (
	"github.com/gin-gonic/gin"

	ctrl "github.com/RodrigoMattosoSilveira/gradeit/controllers/assignment"
	srvc "github.com/RodrigoMattosoSilveira/gradeit/services/assignment"
	repo "github.com/RodrigoMattosoSilveira/gradeit/repository/assignment"
)

// TODO move it into the configs directory ... I get an import cycle error
func RoutesAssignmentt(router *gin.Engine) {

	// Set up ther person routes
	repo := repo.NewAssignment()
	svc := srvc.AssignmentSvcInt(repo)
	routes := ctrl.NewAssignment(svc)


	router.POST("/assignment", routes.Create)
	router.GET("/assignment", routes.GetAll)
	router.GET("/assignment/:id", routes.GetByID)
	router.PUT("/assignment/:id", routes.Update)
	router.DELETE("/assignment/:id", routes.Delete)
}