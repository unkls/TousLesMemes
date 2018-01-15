package actions

import (

  "github.com/jinzhu/gorm"
  "github.com/gorilla/mux"
  "github.com/sirupsen/logrus"
  "net/http"
  _"github.com/go-sql-driver/mysql"

)

type App struct {
  Db      *gorm.DB
  Router  *mux.Router
}

func (a *App) InitializeDatabase() {

  var err error

  a.Db, err = gorm.Open("mysql", "roskva:thialfi@tcp(127.0.0.1:3306)/alpheratz?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    logrus.Fatal(err.Error())
    a.Db.Close()
  }
  a.Db.LogMode(true)

  logrus.Info("Database initialized")

}

func (a *App) InitializeRoutes() {

  a.Router = mux.NewRouter().StrictSlash(true)

  //a.Router.Handle("/token", Use(setHeader).ThenFunc(a.)).Methods("POST")
  a.Router.Handle("/users", Use(log, setHeader).ThenFunc(a.createUser)).Methods("POST")
  a.Router.Handle("/users", Use(log, setHeader).ThenFunc(a.showUsers)).Methods("GET")
  a.Router.Handle("/users/{id}", Use(log, setHeader).ThenFunc(a.showUser)).Methods("GET")
  a.Router.Handle("/users/{id}", Use(log, setHeader).ThenFunc(a.deleteUser)).Methods("DELETE")
  a.Router.Handle("/users/{id}", Use(log, setHeader).ThenFunc(a.updateUser)).Methods("PUT")

  logrus.Info("Routes initialized")

}

func (a *App) Run() {
  logrus.Info("Running on port 8080")
  logrus.Fatal(http.ListenAndServe(":8080", a.Router))
}
