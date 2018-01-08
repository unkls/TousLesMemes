package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/unkls/touslesmemes/database"
)

//InsertUser insert user in the database (receive json user wich contain all the data to insert)
func InsertUser(c echo.Context) error {

	userInsert := &user{}
	stmt, err := database.DbClient.Prepare("INSERT INTO USERS VALUES (?, ?, ?, ?, ?)") // id pseudo email pwd_id
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}
	res, err := stmt.Exec(userInsert.Nom, userInsert.Prenom, userInsert.Email, userInsert.Age)
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}
	id, _ := res.LastInsertId()

	return c.JSON(http.StatusCreated, id)
}
