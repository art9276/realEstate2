package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// User represents a user.
type User struct {
	Id_user       int    `json:"Id_user"`
	Name          string `json:"Name"`
	Surename      string `json:"Surename"`
	Login         string `json:"Login"`
	Enc_password  string `json:"Enc_password"`
	Telephone     string `json:"Telephone"`
	Email         string `json:"Email"`
	Date_creation string `json:"Date_creation"`
	Role          string `json:"Role"`
}

type Users struct {
	Users []User `json:"Users"`
}

func (u User) TableName() string {
	// custom table name, this is default
	return "Users"
}

func ValidateUserInsert(u *User) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(5, 15)),
		validation.Field(&u.Surename, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Login, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Enc_password, validation.Required, validation.Length(3, 50)),
		validation.Field(&u.Telephone, validation.Required, validation.Length(5, 30)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Date_creation, validation.Required, validation.Length(5, 30)),
		validation.Field(&u.Role, validation.Required, validation.Length(4, 12)))
	return err
}

func ValidateUserUpdate(u *User) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Length(5, 15)),
		validation.Field(&u.Surename, validation.Length(5, 20)),
		validation.Field(&u.Login, validation.Length(5, 20)),
		validation.Field(&u.Enc_password, validation.Length(3, 50)),
		validation.Field(&u.Telephone, validation.Length(5, 30)),
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.Date_creation, validation.Length(5, 30)),
		validation.Field(&u.Role, validation.Length(4, 12)))
	return err
}
