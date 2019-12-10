package auth

import (
  "github.com/jc3m/ridge/database"
)

func authenticate(email string, password string) bool {
  db, err := database.Connect()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // Attempt to fetch user from table
  stmt, err := db.Prepare(
    "SELECT password_salt, password_hash FROM users WHERE email = ?")
  if err != nil {
    panic(err)
  }
  defer stmt.Close()
  rows, err := stmt.Query(email)
  if err != nil {
    panic(err)
  }
  if !rows.Next() {
    return false;
  }
  var salt, hash string
  err = rows.Scan(&salt, &hash)
  if err != nil {
    panic(err)
  }
  // Check password
  return hash == genHash(password, salt)
}
