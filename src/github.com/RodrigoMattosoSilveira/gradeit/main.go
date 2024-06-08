package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	cfg "github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"

	ctrlPerson "github.com/RodrigoMattosoSilveira/gradeit/controllers/person"
	srvcPerson "github.com/RodrigoMattosoSilveira/gradeit/services/person"
	repoPerson "github.com/RodrigoMattosoSilveira/gradeit/repository/person"

	ctrlAssignment "github.com/RodrigoMattosoSilveira/gradeit/controllers/assignment"
	srvcAssignment "github.com/RodrigoMattosoSilveira/gradeit/services/assignment"
	repoAssignment "github.com/RodrigoMattosoSilveira/gradeit/repository/assignment"

)

func GetRouter() *gin.Engine {
	return gin.Default()
}

func main() {
 
	// Configure the service
	cfg.Config()
	cfg.DB.AutoMigrate(&models.Person{})
	cfg.DB.AutoMigrate(&models.Assignment{})

	// initialise gofr object
	router := GetRouter()

	// Set up a testing route
	RoutesPing(router)

	// Set up ther person routes
	RoutesPerson(router)
	RoutesAssignment(router)

	// start the service
	http.ListenAndServe(fmt.Sprintf(":%s",  os.Getenv("HTTP_PORT")), router)
	// 
}

func RoutesPerson(router *gin.Engine) {
	// Set up ther person routes
	repo := repoPerson.NewPerson()
	svc := srvcPerson.NewPerson(repo)
	routes := ctrlPerson.NewPerson(svc)


	router.POST("/person", routes.Create)
	router.GET("/person", routes.GetAll)
	router.GET("/person/:id", routes.GetByID)
	router.PUT("/person/:id", routes.Update)
	router.DELETE("/person/:id", routes.Delete)
}

func RoutesAssignment(router *gin.Engine) {

	// Set up ther person routes
	repo := repoAssignment.NewAssignment()
	svc := srvcAssignment.AssignmentSvcInt(repo)
	routes := ctrlAssignment.NewAssignment(svc)


	router.POST("/assignment", routes.Create)
	router.GET("/assignment", routes.GetAll)
	router.GET("/assignment/:id", routes.GetByID)
	router.PUT("/assignment/:id", routes.Update)
	router.DELETE("/assignment/:id", routes.Delete)
}

func RoutesPing(router *gin.Engine) {

	if router == nil {
		router = gin.Default()
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})	
	})

}