package actions

import (
  "net/http"
  "github.com/google/jsonapi"
  "github.com/sirupsen/logrus"
)

type Middleware func(http.Handler) http.Handler

func Use(mw ...Middleware) Middleware {
    return func(finalAction http.Handler) http.Handler {
        for i := len(mw) - 1; i >= 0; i-- {
            finalAction = mw[i](finalAction)
        }
        return finalAction
    }
}

func (mw Middleware) ThenFunc(finalAction func(http.ResponseWriter, *http.Request)) http.Handler {
    return mw(http.HandlerFunc(finalAction))
}

func setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", jsonapi.MediaType)

		next.ServeHTTP(w, r)
	})
}

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    log := logrus.WithFields(logrus.Fields{
      "uri": r.URL.Path,
      "methods": r.Method,
      "host": r.RemoteAddr,
    })

    ctx := contextWithLog(r.Context(), log)

    next.ServeHTTP(w, r.WithContext(ctx))

  })
}
