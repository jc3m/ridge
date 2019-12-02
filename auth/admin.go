package auth

import (
  "errors"

  "github.com/jc3m/ridge/database"
)

func CreateInitialAdmin(email string, password string) error {
  // TODO: Validate email and password
  db, err := database.Connect()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // Check that no other users exist
  row := db.QueryRow("SELECT COUNT(*) FROM users;")
  var userCount int
  err = row.Scan(&userCount)
  if err != nil {
    panic(err)
  }
  if userCount != 0 {
    return errors.New("Users already exist, will not create a new admin account")
  }

  stmt, err := db.Prepare(
    "INSERT INTO users(email, password_hash, password_salt, is_admin) VALUES(?, ?, ?, ?)")
	if err != nil {
    panic(err)
	}
  defer stmt.Close()
  
  salt := genSalt()
  hash := genHash(password, salt)
  _, err = stmt.Exec(email, hash, salt, 1)
  if err != nil {
    panic(err)
  }
  return nil 
}
