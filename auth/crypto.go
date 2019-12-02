package auth

import (
  "crypto/sha256"
  "crypto/rand"
  "encoding/hex"
)

/**
 * Generates a 24 character random salt
 */
func genSalt() string {
  b := make([]byte, 12)
  _, err := rand.Read(b)
  if err != nil {
    panic(err)
  }
  return hex.EncodeToString(b)
}

func genHash(password string, salt string) string {
  var sum [32]byte = sha256.Sum256([]byte(password + salt))
  return hex.EncodeToString(sum[:])
}