package model

type PlayerFilters struct {
	PlayerName  string `json:"player_name"`
	Index       int    `json:"index"`
	Nationality string `json:"nationality"`
}

type ClubFilters struct {
	ClubName string `json:"club_name"`
	Index    int    `json:"index"`
}
