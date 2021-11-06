package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db := SetupDb()

	queryStore := LoadQueriesFromDisk()

	dbStore := DbContext{DB: db, Queries: queryStore}

	r := gin.Default()

	r.Use(RegisterDbContext(dbStore))

	r.LoadHTMLGlob("templates/**/*")

	ConfigureRoutes(*r)

	r.Run("0.0.0.0:4000")
}
