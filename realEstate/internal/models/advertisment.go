package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Advertisment struct {
	IdAdvertisment     int    `json:"IdAdvertisment"`
	TypeAdvertisment   string `json:"TypeAdvertisment"`
	Price              int    `json:"Price"`
	TotalArea          int    `json:"TotalArea"`
	YearOfContribution int    `json:"YearOfContribution"`
	Address            string `json:"Address"`
	Description        string `json:"Description"`
	NumberOfRooms      int    `json:"NumberOfRooms"`
	IsCommercial       int    `json:"IsCommercial"`
}

func ValidateAdvertismentInsert(a *Advertisment) error {
	err := validation.ValidateStruct(a,
		validation.Field(&a.IdAdvertisment, validation.Required, validation.Length(1, 6)),
		validation.Field(&a.TypeAdvertisment, validation.Required, validation.Length(5, 20)),
		validation.Field(&a.Price, validation.Required, validation.Length(5, 20)),
		validation.Field(&a.TotalArea, validation.Required, validation.Length(3, 10)),
		validation.Field(&a.YearOfContribution, validation.Required, validation.Length(4, 4)),
		validation.Field(&a.Address, validation.Required, validation.Length(10, 25)),
		validation.Field(&a.Description, validation.Required, validation.Length(5, 1000)),
		validation.Field(&a.NumberOfRooms, validation.Required, validation.Length(1, 3)),
		validation.Field(&a.IsCommercial, validation.Required, validation.Length(1, 1)))
	return err
}

func ValidateAdvertismentUpdate(a *Advertisment) error {
	err := validation.ValidateStruct(a,
		validation.Field(&a.IdAdvertisment, validation.Length(1, 6)),
		validation.Field(&a.TypeAdvertisment, validation.Length(5, 20)),
		validation.Field(&a.Price, validation.Length(5, 20)),
		validation.Field(&a.TotalArea, validation.Length(3, 10)),
		validation.Field(&a.YearOfContribution, validation.Length(4, 4)),
		validation.Field(&a.Address, validation.Length(10, 25)),
		validation.Field(&a.Description, validation.Length(5, 1000)),
		validation.Field(&a.NumberOfRooms, validation.Length(1, 3)),
		validation.Field(&a.IsCommercial, validation.Length(1, 1)))
	return err
}
