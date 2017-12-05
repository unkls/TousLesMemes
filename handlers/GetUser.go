package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/unkls/touslesmemes/database"
)

type user struct {
	Nom    string `json:"private"`
	Prenom string `json:"public"`
	Pseudo string `json:"pseudo"`
	Email  string `json:"email"`
	Age    string `json:"age"`
}

// GetUser return the user with the parameter ID
func GetUser(c echo.Context) error {
	idUser := c.Param("id")
	var databaseUser string
	err := database.DbClient.QueryRow("SELECT * FROM USERS WHERE user_id = ?", idUser).Scan(&databaseUser)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return c.JSON(http.StatusOK, databaseUser)
}
