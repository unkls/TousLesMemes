package main

import (
	"github.com/labstack/echo"
	"github.com/unkls/touslesmemes/config"
	"github.com/unkls/touslesmemes/handlers"
)

func main() {
	config.LoadConfig()
	e := echo.New()

	e.GET("/status", handlers.Status)
	e.GET("/GetUser/:id", handlers.GetUser)
	e.GET("/GetPassword", handlers.GetPassword)
	e.POST("/InsertUser", handlers.InsertUser)
	e.Logger.Fatal(e.Start(":8080"))
}
