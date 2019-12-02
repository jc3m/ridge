package auth

import (
  "fmt"
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
)

func AuthRouter(s *mux.Router) {
  s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "auth")
  })

  s.HandleFunc("/createadmin", func(w http.ResponseWriter, r *http.Request) {
    type createAdminRequest struct {
      Email string;
      Password string;
    }

    var t createAdminRequest
    err := json.NewDecoder(r.Body).Decode(&t)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return;
    }

    err = CreateInitialAdmin(t.Email, t.Password)
    if err != nil {
      w.WriteHeader(http.StatusForbidden)
      fmt.Fprintf(w, err.Error())
    }
  }).Methods("POST")
}
