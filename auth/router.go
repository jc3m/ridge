package auth

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func AuthRouter(s *mux.Router) {
  s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "auth")
  })
}
