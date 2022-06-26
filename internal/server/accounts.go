package server

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	//Posts      []*Post
}

var Accounts []*Account

func AddAccount(account *Account) error {

	encrypt_acc := EncryptPassword(account)
	db := GetDBConnection()
	_, err := db.Model(encrypt_acc).Returning("*").Insert()
	if err != nil {
		return err
	}
	return nil
}
func AddPost(account *Account, post *Post) {
	//TODO: Add post for each
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

func EncryptPassword(account *Account) *Account {

	new_pass, error := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if error != nil {
		return account
	}

	account.Password = string(new_pass)
	return account
}
