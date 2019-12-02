package auth

import (
  "crypto/sha256"
  "crypto/rand"
  "encoding/hex"
  "errors"

  "github.com/jc3m/ridge/database"
)

// TODO: Move these to a crypto specific file

/**
 * Generates a 24 character random salt
 */
func genSalt() (string, error) {
  b := make([]byte, 12)
  _, err := rand.Read(b)
  if err != nil {
    return "", err
  }
  return hex.EncodeToString(b), nil
}

func genHash(password string, salt string) string {
  var sum [32]byte = sha256.Sum256([]byte(password + salt))
  return hex.EncodeToString(sum[:])
}

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
    "INSERT INTO users(email, password_hash, password_salt, is_admin) VALUES( ?, ?, ?, ? )")
	if err != nil {
    panic(err)
	}
  defer stmt.Close()
  
  salt, err := genSalt()
  if err != nil {
    panic(err)
  }
  hash := genHash(password, salt)
  _, err = stmt.Exec(email, hash, salt, 1)
  if err != nil {
    panic(err)
  }
  return nil 
}
