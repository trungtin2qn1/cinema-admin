package utils

import "golang.org/x/crypto/bcrypt"

//Generate a salted hash for the input string
func Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare ...
func Compare(hash string, s string) (bool, error) {
	incoming := []byte(s)
	existing := []byte(hash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	if err != nil {
		return false, err
	}
	return true, nil
}
