package api

import (
  "net/http"

  "github.com/gorilla/mux"
  "github.com/jc3m/ridge/auth"
)

func Router() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  auth.AuthRouter(r.PathPrefix("/auth").Subrouter())

  return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
  Router().ServeHTTP(w, r)
}
