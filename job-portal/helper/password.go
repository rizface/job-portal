package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plain string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(plain),bcrypt.DefaultCost)
	return string(hash),err
}

func ComparePassword(plain string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(plain))
	return err
}