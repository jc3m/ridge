package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/joho/godotenv"

  "github.com/jc3m/ridge/auth"
)

func main() {
  // Loads from .env file. Fails silently if .env doesn't exist
  godotenv.Load()

  r := mux.NewRouter().StrictSlash(true)
  auth.AuthRouter(r.PathPrefix("/auth").Subrouter())

  // TODO: Restrict to https
  http.Handle("/", r)
  // TODO: Read port from flag
  log.Fatal(http.ListenAndServe(":8080", nil))
}
