package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserId   int
	Alias    string
	FullName string
	Active   bool
	Locked   bool
}

func GetAllUsersJson(c *gin.Context) {
	db := GetDb(c)

	rows, err := db.Query("Select * From Users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Query Error",
		})
	}
	defer rows.Close()

	var users []UserModel
	for rows.Next() {
		var user UserModel
		if err := rows.Scan(&user.UserId, &user.Alias,
			&user.FullName, &user.Active, &user.Locked); err != nil {
			if len(users) == 0 {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "Error reading user",
				})
			}
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
