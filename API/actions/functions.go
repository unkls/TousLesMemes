package actions

import (

  "net/http"
  "github.com/gorilla/mux"
  "context"
  "touslesmemes/ZAPI/actions/models"
  "github.com/sirupsen/logrus"
)

func logFromContext(ctx context.Context) *logrus.Entry {
  return ctx.Value(logKey).(*logrus.Entry)
}

func userFromContext(ctx context.Context) *models.User {
  return ctx.Value(usrKey).(*models.User)
}

func contextWithUser(ctx context.Context, user *models.User) context.Context {
  return context.WithValue(ctx, usrKey, user)
}

func contextWithLog(ctx context.Context, log *logrus.Entry) context.Context {
  return context.WithValue(ctx, logKey, log)
}

func fetchVar(r *http.Request, variable string) string {
  return mux.Vars(r)[variable]
}

func exitOK(w http.ResponseWriter, code int, detail string, log *logrus.Entry) {

  //w.WriteHeader(code)
  log.Info(detail)

}

func exitWithError(w http.ResponseWriter, detail string, code int, log *logrus.Entry) {

  log.Error(detail)
  http.Error(w, detail, code)
  return

}
