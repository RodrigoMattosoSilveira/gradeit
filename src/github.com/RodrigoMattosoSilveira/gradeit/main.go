package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	cfg "github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/models"
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
