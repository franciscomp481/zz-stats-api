package webscrapper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/franciscomp481/zerozero-stats-api/model"
)

func FetchClubStats(doc *goquery.Document) (model.ClubStats, error) {
	// Extract the club stats from the HTML content
	var clubName string
	var season string
	var market_value string
	var stats model.ResultStats

	doc.Find("div.bio").Each(func(i int, s *goquery.Selection) {
		span := s.Find("span").Text()
		if span == "Nome" {
			clubName = s.Contents().Not("span").Text()
		}
	})

	doc.Find("form[name='form_equipa']").Each(func(i int, s *goquery.Selection) {
		// Find the select element with the name epoca_id
		s.Find("select[name='epoca_id']").Each(func(i int, selectElem *goquery.Selection) {
			// Find the selected option within the select element
			selectElem.Find("option[selected]").Each(func(i int, option *goquery.Selection) {
				season = option.Text()
			})
		})
	})

	doc.Find("div.rectangle").Each(func(i int, s *goquery.Selection) {
		s.Find("div.value").Each(func(i int, value *goquery.Selection) {
			market_value = value.Text()
		})
	})

	// Initialize the map to store result stats per competition
	resultStatsPerCompetition := make(map[string]model.ResultStats)

	// Iterate through each row in the table that represents either a competition's stats or the totals
	doc.Find("div.box h2.header:contains('Resumo') + div.box_table table.zztable.stats tbody tr").Each(func(i int, s *goquery.Selection) {
		// Check if the row contains totals by looking for 'td.totals'
		tds := s.Find("td.totals")

		if tds.Length() >= 6 {
			// This is the totals row
			matchesPlayed, err := strconv.Atoi(strings.TrimSpace(tds.Eq(1).Text()))
			if err != nil {
				fmt.Println("Error converting MatchesPlayed:", err)
			}
			victories, err := strconv.Atoi(strings.TrimSpace(tds.Eq(2).Text()))
			if err != nil {
				fmt.Println("Error converting Victories:", err)
			}
			draws, err := strconv.Atoi(strings.TrimSpace(tds.Eq(3).Text()))
			if err != nil {
				fmt.Println("Error converting Draws:", err)
			}
			defeats, err := strconv.Atoi(strings.TrimSpace(tds.Eq(4).Text()))
			if err != nil {
				fmt.Println("Error converting Defeats:", err)
			}

			// Split the goals data by the "-" character
			goals := strings.Split(strings.TrimSpace(tds.Eq(5).Text()), "-")
			if len(goals) == 2 {
				goalsScored, err := strconv.Atoi(goals[0])
				if err != nil {
					fmt.Println("Error converting GoalsScored:", err)
				}
				goalsConceded, err := strconv.Atoi(goals[1])
				if err != nil {
					fmt.Println("Error converting GoalsConceded:", err)
				}

				// Populate the ResultStats struct for overall totals
				stats = model.ResultStats{
					MatchesPlayed: matchesPlayed,
					Victories:     victories,
					Draws:         draws,
					Defeats:       defeats,
					GoalsScored:   goalsScored,
					GoalsConceded: goalsConceded,
				}

				// Optionally, store or print the overall totals
				// fmt.Println("Extracted total stats:", stats)
			} else {
				fmt.Println("Unexpected format for goals:", tds.Eq(5).Text())
			}
		} else {
			// Process competition-specific stats
			tds := s.Find("td.stat")

			if tds.Length() >= 5 {
				// Extract the competition name
				competitionName := strings.TrimSpace(s.Find("td.edition div.text a").Text())

				matchesPlayed, err := strconv.Atoi(strings.TrimSpace(tds.Eq(0).Text()))
				if err != nil {
					fmt.Println("Error converting MatchesPlayed:", err)
					return
				}
				victories, err := strconv.Atoi(strings.TrimSpace(tds.Eq(1).Text()))
				if err != nil {
					fmt.Println("Error converting Victories:", err)
					return
				}
				draws, err := strconv.Atoi(strings.TrimSpace(tds.Eq(2).Text()))
				if err != nil {
					fmt.Println("Error converting Draws:", err)
					return
				}
				defeats, err := strconv.Atoi(strings.TrimSpace(tds.Eq(3).Text()))
				if err != nil {
					fmt.Println("Error converting Defeats:", err)
					return
				}

				// Split the goals data by the "-" character
				goals := strings.Split(strings.TrimSpace(tds.Eq(4).Text()), "-")
				if len(goals) == 2 {
					goalsScored, err := strconv.Atoi(goals[0])
					if err != nil {
						fmt.Println("Error converting GoalsScored:", err)
						return
					}
					goalsConceded, err := strconv.Atoi(goals[1])
					if err != nil {
						fmt.Println("Error converting GoalsConceded:", err)
						return
					}

					// Populate the ResultStats struct for this competition
					stats := model.ResultStats{
						MatchesPlayed: matchesPlayed,
						Victories:     victories,
						Draws:         draws,
						Defeats:       defeats,
						GoalsScored:   goalsScored,
						GoalsConceded: goalsConceded,
					}

					// Store the stats in the map under the competition name
					resultStatsPerCompetition[competitionName] = stats
				} else {
					fmt.Println("Unexpected format for goals:", tds.Eq(4).Text())
				}
			}
		}
	})

	//lastGames := make(map[string]model.LastGames)
	NextGames := make([]model.NextGames, 0)

	doc.Find("div.box h2.header:contains('Jogos') + div.box_table tbody tr").Slice(0, 3).Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")

		date := strings.TrimSpace(tds.Eq(1).Text())
		hour := strings.TrimSpace(tds.Eq(2).Text())
		if hour == "" {
			hour = "N/A"
		}
		competition := strings.TrimSpace(tds.Eq(3).Text())
		homeTeam := strings.TrimSpace(tds.Eq(4).Text())
		awayTeam := strings.TrimSpace(tds.Eq(8).Text())

		nextGames := model.NextGames{
			Date:        date,
			Hour:        hour,
			Competition: competition,
			HomeTeam:    homeTeam,
			AwayTeam:    awayTeam,
		}

		NextGames = append(NextGames, nextGames)
	})

	totalRows := doc.Find("div.box h2.header:contains('Jogos') + div.box_table tbody tr").Length()

	LastGames := make([]model.LastGames, 0)

	doc.Find("div.box h2.header:contains('Jogos') + div.box_table tbody tr").Slice(totalRows-3, totalRows).Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")

		form := strings.TrimSpace(tds.Eq(0).Text())
		date := strings.TrimSpace(tds.Eq(1).Text())
		hour := strings.TrimSpace(tds.Eq(2).Text())
		competition := strings.TrimSpace(tds.Eq(3).Text())
		homeTeam := strings.TrimSpace(tds.Eq(4).Text())
		result := strings.TrimSpace(tds.Eq(6).Text())
		awayTeam := strings.TrimSpace(tds.Eq(8).Text())

		lastGames := model.LastGames{
			Form:        form,
			Date:        date,
			Hour:        hour,
			Competition: competition,
			HomeTeam:    homeTeam,
			Result:      result,
			AwayTeam:    awayTeam,
		}

		LastGames = append(LastGames, lastGames)
	})

	clubStats := model.ClubStats{
		TeamName:                  clubName,
		Season:                    season,
		MarketValue:               market_value,
		ResultStats:               stats,
		ResultStatsPerCompetition: resultStatsPerCompetition,
		NextGames:                 NextGames,
		LastGames:                 LastGames,
	}

	return clubStats, nil
}
