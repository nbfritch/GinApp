package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := SetupDb()

	r := gin.Default()

	r.Use(Database(db))

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hello",
		})
	})

	r.Run("0.0.0.0:4000")
}
