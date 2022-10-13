package models

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
