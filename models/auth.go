package models

import (
	"cinema-admin/utils"
	"cinema-admin/utils/jwt"
	"fmt"
)

//AuthReq ...
type AuthReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func checkAuthData(email string, password string) bool {
	if email == "" {
		return false
	}
	if utils.ValidateFormat(email) != nil {
		return false
	}
	if len(password) < 6 {
		return false
	}
	return true
}

//Login ...
func Login(email string, password string) (Consumer, error) {
	consumer := Consumer{}
	var err error
	if !(checkAuthData(email, password)) {
		err = fmt.Errorf("%s", "Email or password is invalid")
		return consumer, err
	}

	consumer, err = GetConsumerByEmail(email)

	if err != nil {
		err = fmt.Errorf("%s", "Consumer is not available")
		return consumer, err
	}

	check, err := utils.Compare(consumer.Password, password)

	if err != nil {
		err = fmt.Errorf("%s", "Password is not right")
		return consumer, err
	}

	if check == false {
		err = fmt.Errorf("%s", "Password is not right")
		return consumer, err
	}

	token, err := jwt.IssueToken(consumer.ID, consumer.Email)
	consumer.Token = token

	return consumer, err
}
