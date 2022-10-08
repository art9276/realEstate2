package password

import (
	"golang.org/x/crypto/bcrypt"
	logg "realEstate/pkg/log"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logg.Error("failed to hash password: %w")
	}
	return string(bytes)
}

// CheckPassword checks if the provided password is correct or not
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
