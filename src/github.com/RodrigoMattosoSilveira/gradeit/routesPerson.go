package main

import (
	"github.com/gin-gonic/gin"

	ctrl "github.com/RodrigoMattosoSilveira/gradeit/controllers/person"
	srvc "github.com/RodrigoMattosoSilveira/gradeit/services/person"
	repo "github.com/RodrigoMattosoSilveira/gradeit/repository/person"
)

// TODO move it into the configs directory ... I get an import cycle error
func RoutesPerson(router *gin.Engine) {
	// Set up ther person routes
	repo := repo.NewPerson()
	svc := srvc.NewPerson(repo)
	routes := ctrl.NewPerson(svc)


	router.POST("/person", routes.Create)
	router.GET("/person", routes.GetAll)
	router.GET("/person/:id", routes.GetByID)
	router.PUT("/person/:id", routes.Update)
	router.DELETE("/person/:id", routes.Delete)
}