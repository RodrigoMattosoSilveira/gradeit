package main

import (
	"fmt"
	"net/http"
	"os"

	cfg "github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/controllers/person"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
	"github.com/RodrigoMattosoSilveira/gradeit/repository"
	"github.com/RodrigoMattosoSilveira/gradeit/services"
	"github.com/gin-gonic/gin"
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
	repo := repository.NewPerson()
	svc := services.NewPerson(repo)
	personRoutes := controllers.NewPerson(svc)

	router.POST("/person", personRoutes.Create)
	router.GET("/person", personRoutes.GetAll)
	router.GET("/person/:id", personRoutes.GetByID)
	router.PUT("/person", personRoutes.Update)
	router.DELETE("/delete", personRoutes.Delete)

	// start the service
	http.ListenAndServe(fmt.Sprintf(":%s",  os.Getenv("HTTP_PORT")), router)
	// route.Run() // listen and serve on 0.0.0.0:8080
}