package models

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
