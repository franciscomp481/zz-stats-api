package webscrapper

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/franciscomp481/zerozero-stats-api/model"
	"golang.org/x/net/html/charset"
)

// Search for the player by name and return the URL

func GetPlayerPage(playerURL string) (*goquery.Document, error) {

	// Fetch the HTML content of the player's profile page
	resp, err := http.Get(playerURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body with the correct encoding
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	return doc, nil

}

func FetchPlayerStats(doc *goquery.Document) (model.PlayerStats, error) {
	playerName := doc.Find("div.zz-enthdr-data span.name").Contents().Not("span.number").Text()
	playerStats := model.PlayerStats{
		PlayerName:  playerName,
		Seasons:     make(map[string][]model.ClubStats),
		Tournaments: make(map[string]model.TournamentStats),
	}

	var lastSeason string

	// Extract data from the player stats table
	doc.Find("div.section:contains('Futebol')").NextAllFiltered("table.career").First().Each(func(i int, table *goquery.Selection) {
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

			// Debugging statements
			//fmt.Printf("Season: %s, Club: %s, Matches Played: %d, Goals Scored: %d, Assists: %d\n",
			//	season, club, matchesPlayed, goalsScored, assists)

			if season != "" {
				clubStats := model.ClubStats{
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
