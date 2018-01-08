package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/unkls/touslesmemes/database"
)

//InsertPassword insert password in the database
func InsertPassword(c echo.Context) error {

	stmt, err := database.DbClient.Prepare("INSERT INTO PASSWORD (PASSWORD.password) VALUES (?)") //pwd_password
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}
	res, err := stmt.Exec("password")
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}
	id, _ := res.LastInsertId()

	return c.JSON(http.StatusCreated, id)
}
