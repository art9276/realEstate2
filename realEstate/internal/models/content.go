package models

type Content struct {
	IdContent    int    `json:"IdContent"`
	Article      string `json:"Article"`
	DateCreation string `json:"DateCreation"`
	AuthorID     int    `json:"AuthorID"`
	Text         string `json:"Text"`
	MiniContent  string `json:"MiniContent"`
}
