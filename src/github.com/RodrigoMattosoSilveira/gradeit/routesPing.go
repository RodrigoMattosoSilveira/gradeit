package main

import (
	"github.com/gin-gonic/gin"
)

// TODO move it into the configs directory ... I get an import cycle error
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