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

func RegisterUserPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create_user.page.html", nil)
}

func RegisterUserPostHandler(c *gin.Context) {
	alias := c.PostForm("alias")
	if len(alias) == 0 {
		c.HTML(http.StatusBadRequest, "create_user.page.html", gin.H{
			"ErrorMessage": "`Alias` is required",
		})
	}

	fullName := c.PostForm("fullName")
	if len(fullName) == 0 {
		c.HTML(http.StatusBadRequest, "create_user.page.html", gin.H{
			"ErrorMessage": "`Full Name` is required",
		})
	}

	c.Redirect(http.StatusCreated, "")
}

func GetAllUsersJson(c *gin.Context) {
	db := GetDbContext(c)

	rows, err := db.DB.Query("Select * From Users")

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
