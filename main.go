package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db := SetupDb()

	r := gin.Default()

	r.Use(Database(db))

	r.LoadHTMLGlob("templates/**/*")

	ConfigureRoutes(*r)

	r.Run("0.0.0.0:4000")
}
