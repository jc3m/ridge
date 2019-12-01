package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
)

// TODO: Read from flag
var port string = "8000"

func main() {
  // Loads from .env file. Fails silently if .env doesn't exist
  godotenv.Load()

  r := mux.NewRouter()

  // TODO: Restrict to https
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
