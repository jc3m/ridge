package auth

import (
  "fmt"
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
)

type userRequest struct {
  Email string;
  Password string;
}

func AuthRouter(s *mux.Router) {
  s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "auth")
  })

  s.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    var u userRequest
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return;
    }

    if !authenticate(u.Email, u.Password) {
      w.WriteHeader(http.StatusUnauthorized)
      return;
    }
    // TODO: Create and return a session
  }).Methods("POST")

  s.HandleFunc("/createadmin", func(w http.ResponseWriter, r *http.Request) {
    var u userRequest
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return;
    }

    err = createInitialAdmin(u.Email, u.Password)
    if err != nil {
      w.WriteHeader(http.StatusForbidden)
      fmt.Fprintf(w, err.Error())
    }
  }).Methods("POST")
}
