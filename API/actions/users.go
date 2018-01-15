package actions

import (

  "net/http"

  "touslesmemes/ZAPI/actions/models"
  "github.com/google/jsonapi"
  "github.com/satori/uuid"

)

func (a *App) me(w http.ResponseWriter, r *http.Request) {

  log := logFromContext(r.Context())
  u := userFromContext(r.Context())

  if err := jsonapi.MarshalPayload(w, u); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  exitOK(w, http.StatusNoContent, "OK", log)

}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {

  log := logFromContext(r.Context())
  u := new(models.User)
  u.ID = fetchVar(r, "id")

  if err := u.Delete(a.Db); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  exitOK(w, http.StatusNoContent, "OK", log)

}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {

  log := logFromContext(r.Context())
  u := new(models.User)

  if err := jsonapi.UnmarshalPayload(r.Body, u); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  u.ID = uuid.Must(uuid.NewV4()).String()

  if err := u.Create(a.Db); err!= nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  if err := jsonapi.MarshalPayload(w, u); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  exitOK(w, http.StatusNoContent, "OK", log)

}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request)   {

  log := logFromContext(r.Context())
  u := new(models.User)
  update := new(models.User)
  u.ID = fetchVar(r, "id")

  if err := jsonapi.UnmarshalPayload(r.Body, update); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }


  if err := u.Update(a.Db, update); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  u.HidePassword()

  if err := jsonapi.MarshalPayload(w, u); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  exitOK(w, http.StatusNoContent, "OK", log)

}

func (a *App) showUser(w http.ResponseWriter, r *http.Request) {

  log := logFromContext(r.Context())
  u := new(models.User)
  u.ID = fetchVar(r, "id")

  if err := u.Show(a.Db); err!= nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  u.HidePassword()

  if err := jsonapi.MarshalPayload(w, u); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  exitOK(w, http.StatusNoContent, "OK", log)

}

func (a *App) showUsers(w http.ResponseWriter, r *http.Request) {

  log := logFromContext(r.Context())
  u := models.Users{}

  if err := u.Show(a.Db); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  u.HidePassword()

  if err := jsonapi.MarshalPayload(w, u); err != nil {
    exitWithError(w, err.Error(), http.StatusInternalServerError, log)
    return
  }

  exitOK(w, http.StatusNoContent, "OK", log)

}
