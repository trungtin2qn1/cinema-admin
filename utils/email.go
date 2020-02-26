package utils

import (
	"errors"
	"regexp"
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// EmailValidate ...
func EmailValidate(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("Email is invalid")
	}
	return nil
}
