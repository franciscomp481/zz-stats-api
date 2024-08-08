package model

type SearchPlayerQuery struct {
	PlayerName string `json:"player_name"`
	Index      int    `json:"index"`
}
