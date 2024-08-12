package model

type PlayerFilters struct {
	PlayerName  string `json:"player_name"`
	Index       int    `json:"index"`
	Nationality string `json:"nationality"`
}
