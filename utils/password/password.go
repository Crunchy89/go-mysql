package password

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	data := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash[:])
}

func VerifyPassword(password, hash string) bool {
	pass := []byte(password)
	hashed := []byte(hash)
	if err := bcrypt.CompareHashAndPassword(hashed, pass); err != nil {
		return false
	}
	// nil means it is a match
	return true
}
