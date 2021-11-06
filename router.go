package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AliveHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "api is alive",
	})
}

func ConfigureRoutes(r gin.Engine) {
	// Index
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.page.html", nil)
	})

	users := r.Group("/users")
	{
		users.GET("/register", RegisterUserPageHandler)
	}

	// API group
	api := r.Group("/api")
	{
		api.GET("/", AliveHandler)
		api.GET("/users", GetAllUsersJson)
	}
}
