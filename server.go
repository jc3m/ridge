package main

import (
  "flag"
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"

  "github.com/jc3m/ridge/auth"
)

var port = flag.Int("port", 8080, "port to listen for requests")

func main() {
  // Loads from .env file. Fails silently if .env doesn't exist
  godotenv.Load()

  r := mux.NewRouter().StrictSlash(true)
  auth.AuthRouter(r.PathPrefix("/auth").Subrouter())

  r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

  // TODO: Restrict to https
  http.Handle("/", r)

  log.Printf("Listening on port %d\n", *port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
