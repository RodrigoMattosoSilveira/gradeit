package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	cfg "github.com/RodrigoMattosoSilveira/gradeit/configs"
	ctrlPerson "github.com/RodrigoMattosoSilveira/gradeit/controllers/person"
	svcPerson "github.com/RodrigoMattosoSilveira/gradeit/services/person"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	repoPerson "github.com/RodrigoMattosoSilveira/gradeit/repository/person"
)

func main() {
 
	// Configure the service
	cfg.Config()
	cfg.DB.AutoMigrate(&models.Person{})

	// initialise gofr object
	router := gin.Default()

	// Set up a testing route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Set up ther person routes
	repo := repoPerson.NewPerson()
	svc := svcPerson.NewPerson(repo)
	personRoutes := ctrlPerson.NewPerson(svc)

	router.POST("/person", personRoutes.Create)
	router.GET("/person", personRoutes.GetAll)
	router.GET("/person/:id", personRoutes.GetByID)
	router.PUT("/person/:id", personRoutes.Update)
	router.DELETE("/person/:id", personRoutes.Delete)

	// start the service
	http.ListenAndServe(fmt.Sprintf(":%s",  os.Getenv("HTTP_PORT")), router)
	// 
}
	
	func personRoutes(router *gin.Engine) {

		// Set up ther person routes
		repo := repoPerson.NewPerson()
		svc := svcPerson.NewPerson(repo)
		personRoutes := ctrlPerson.NewPerson(svc)
	
	
		router.POST("/person", personRoutes.Create)
		router.GET("/person", personRoutes.GetAll)
		router.GET("/person/:id", personRoutes.GetByID)
		router.PUT("/person/:id", personRoutes.Update)
		router.DELETE("/person/:id", personRoutes.Delete)
	
	}