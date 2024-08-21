package model

type PlayerClubStats struct {
	Club          string `json:"club"`
	MatchesPlayed int    `json:"matches_played"`
	GoalsScored   int    `json:"goals_scored"`
	Assists       int    `json:"assists"`
}

type PlayerStats struct {
	PlayerName  string                       `json:"player_name"`
	Seasons     map[string][]PlayerClubStats `json:"seasons"`
	Tournaments map[string]TournamentStats   `json:"tournaments"`
}

type TournamentStats struct {
	MatchesPlayed int `json:"matches_played"`
	GoalsScored   int `json:"goals_scored"`
	Assists       int `json:"assists"`
}
