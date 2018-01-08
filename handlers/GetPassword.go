package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/unkls/touslesmemes/database"
)

// GetPassword return the password of one user
func GetPassword(c echo.Context) error {
	id := c.Param("id")
	var databasePassword string
	err := database.DbClient.QueryRow("SELECT USERS.password FROM PASSWORD WHERE user_id = ?", id).Scan(&databasePassword)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return c.JSON(http.StatusOK, databasePassword)
}
