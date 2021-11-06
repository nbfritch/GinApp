package main

import (
	"github.com/gin-gonic/gin"
)

const DB_KEY = "database"

func RegisterDbContext(d DbContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(DB_KEY, d)
	}
}

func GetDbContext(c *gin.Context) DbContext {
	return c.MustGet(DB_KEY).(DbContext)
}
