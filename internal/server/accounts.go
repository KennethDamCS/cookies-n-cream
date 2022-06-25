package server

import (
	"errors"
)

type Account struct {
	Email    string `json:"email" binding:"required,email"`
	Password string
}

var Accounts []*Account

func AddAccount(account *Account) error {
	db := GetDBConnection()
	_, err := db.Model(account).Returning("*").Insert()
	if err != nil {
		return err
	}
	return nil
}

func Authenticate(email, password string) (*Account, error) {
	db := GetDBConnection()
	account := new(Account)
	if err := db.Model(account).Where(
		"email = ?", email).Select(); err != nil {
		return nil, err
	}
	if password != account.Password {
		return nil, errors.New("Password not valid.")
	}
	return account, nil
}
