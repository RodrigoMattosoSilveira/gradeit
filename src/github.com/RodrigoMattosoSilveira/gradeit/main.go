package main

import (
	"fmt"
	"net/http"
	"os"

	"log/slog"

	cfg "github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/controllers/person"
	"github.com/RodrigoMattosoSilveira/gradeit/repository"
	"github.com/RodrigoMattosoSilveira/gradeit/services"
	"github.com/gin-gonic/gin"
)

func main() {
    // initialise gofr object
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := cfg.SetEnv()
	if err != nil {
		panic("Error setting the environment")
	}

	// Showcase environment variables handling
	// without setting the environment
	//
	slog.Info(fmt.Sprintf("THIS_ENV = %s",  os.Getenv("THIS_ENV")))
	slog.Info(fmt.Sprintf("HTTP_PORT = %s",  os.Getenv("HTTP_PORT")))
	slog.Info(fmt.Sprintf("DB_DIALECT = %s",  os.Getenv("DB_DIALECT")))
	slog.Info(fmt.Sprintf("DB_NAME = %s",  os.Getenv("DB_NAME")))

	repo := repository.NewPerson()
	svc := services.NewPerson(repo)
	personRoutes := controllers.NewPerson(svc)

	router.POST("/person", personRoutes.Create)
	router.GET("/person", personRoutes.GetAll)
	router.GET("/person/:id", personRoutes.GetByID)
	router.PUT("/person", personRoutes.Update)
	router.DELETE("/delete", personRoutes.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s",  os.Getenv("HTTP_PORT")), router)
	// route.Run() // listen and serve on 0.0.0.0:8080
}