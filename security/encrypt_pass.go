package security

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckEncryptPassword(password string, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
