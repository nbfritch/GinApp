package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

const DB_KEY = "database"

func Database(d *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(DB_KEY, d)
	}
}

func GetDb(c *gin.Context) *sql.DB {
	d := c.MustGet(DB_KEY).(sql.DB)
	return &d
}
