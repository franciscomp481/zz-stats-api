package webscrapper

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/franciscomp481/zerozero-stats-api/model"
)

func FetchPlayerStats(doc *goquery.Document) (model.PlayerStats, error) {
	playerName := doc.Find("div.zz-enthdr-data span.name").Contents().Not("span.number").Text()
	playerStats := model.PlayerStats{
		PlayerName:  playerName,
		Seasons:     make(map[string][]model.PlayerClubStats),
		Tournaments: make(map[string]model.TournamentStats),
	}

	var lastSeason string

	// Extract data from the player stats table
	doc.Find("div.section").FilterFunction(func(i int, s *goquery.Selection) bool {
		sectionText := s.Text()
		return strings.Contains(sectionText, "Futsal") || strings.Contains(sectionText, "Futebol")
	}).NextAllFiltered("table.career").First().Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			var season, club string
			var matchesPlayed, goalsScored, assists int

			row.Find("td").Each(func(k int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text()) // Trim whitespace
				switch k {
				case 0:
					// skip
				case 1:
					season = text
					if season != "" {
						lastSeason = season
					}
				case 2:
					club = text
				case 3:
					matchesPlayed, _ = strconv.Atoi(text)
				case 4:
					goalsScored, _ = strconv.Atoi(text)
				case 5:
					assists, _ = strconv.Atoi(text)
				}
			})

			// Use last non-empty season if current season is empty
			if season == "" {
				season = lastSeason
			}

			if season != "" {
				clubStats := model.PlayerClubStats{
					Club:          club,
					MatchesPlayed: matchesPlayed,
					GoalsScored:   goalsScored,
					Assists:       assists,
				}
				playerStats.Seasons[season] = append(playerStats.Seasons[season], clubStats)
			}
		})
	})

	doc.Find("div.section:contains('EDIÇÕES')").NextAllFiltered("table.career").First().Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			var tournament_name string
			var matchesPlayed, goalsScored, assists int

			row.Find("td").Each(func(k int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text()) // Trim whitespace
				switch k {
				case 0:
					// skip
				case 1:
					// skip
				case 2:
					tournament_name = text
				case 3:
					matchesPlayed, _ = strconv.Atoi(text)
				case 4:
					goalsScored, _ = strconv.Atoi(text)
				case 5:
					assists, _ = strconv.Atoi(text)
				}
			})

			if tournament_name != "" {
				TournamentStats := model.TournamentStats{
					MatchesPlayed: matchesPlayed,
					GoalsScored:   goalsScored,
					Assists:       assists,
				}
				playerStats.Tournaments[tournament_name] = TournamentStats
			}
		})
	})

	return playerStats, nil
}
