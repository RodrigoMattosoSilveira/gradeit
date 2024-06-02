package main

import (
	"fmt"
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"
	c "github.com/RodrigoMattosoSilveira/gradeit/configs"
)

func main() {
    // initialise gofr object
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := c.SetEnv()
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

	r.Run() // listen and serve on 0.0.0.0:8080
}