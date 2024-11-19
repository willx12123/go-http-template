package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func checkPassword(pwdDigest, pwd string) bool {
	if pwdDigest == "" && pwd != "" {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(pwdDigest), []byte(pwd)) == nil
}

func encodePassword(pwd string) (string, error) {
	if pwd == "" {
		return "", errors.New("password can not be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generate password hash: %w", err)
	}

	return string(hash), nil
}
