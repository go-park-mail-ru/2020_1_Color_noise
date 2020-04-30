package utils

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(encryptedPassword string, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil {
		return nil
	}

	return BadPassword.Newf("Password is incorrect")
}

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", Wrapf(err, "Error in during encrypting password")
	}

	return string(hash), nil
}
