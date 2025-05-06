package model

import "net/mail"

type User struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) IsValidEmail() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}
