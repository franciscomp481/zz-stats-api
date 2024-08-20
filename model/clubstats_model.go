package model

type ClubStats struct {
	TeamName                  string                 `json:"team_name"`
	Season                    string                 `json:"season"`
	ResultStats               ResultStats            `json:"result_stats"`
	ResultStatsPerCompetition map[string]ResultStats `json:"result_stats_per_competition"`
	MarketValue               string                 `json:"market_value"`
	LastGames                 map[string]LastGames   `json:"last_games"`
	NextGames                 map[string]NextGames   `json:"next_games"`
}

type NextGames struct {
	Date        string `json:"date"`
	Hour        string `json:"hour"`
	Competition string `json:"competition"`
	HomeTeam    string `json:"home_team"`
	AwayTeam    string `json:"away_team"`
}

type LastGames struct {
	Form        string `json:"form"`
	Date        string `json:"date"`
	Hour        string `json:"hour"`
	Competition string `json:"competition"`
	HomeTeam    string `json:"home_team"`
	Result      string `json:"result"`
	AwayTeam    string `json:"away_team"`
}

type ResultStats struct {
	MatchesPlayed int `json:"matches_played"`
	Victories     int `json:"victories"`
	Draws         int `json:"draws"`
	Defeats       int `json:"defeats"`
	GoalsScored   int `json:"goals_scored"`
	GoalsConceded int `json:"goals_conceded"`
}
