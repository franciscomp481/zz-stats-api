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

	doc.Find("div.box h2.header:contains('Resumo') + div.box_table table.zztable.stats tbody tr").Each(func(i int, s *goquery.Selection) {
		// Find all 'td' elements within this row
		tds := s.Find("td.totals")

		// Ensure there are enough 'td' elements
		if tds.Length() >= 6 {
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

				// Populate the ResultStats struct
				stats = model.ResultStats{
					MatchesPlayed: matchesPlayed,
					Victories:     victories,
					Draws:         draws,
					Defeats:       defeats,
					GoalsScored:   goalsScored,
					GoalsConceded: goalsConceded,
				}

				// Store the stats in the map

				// fmt.Println("Extracted stats:", stats)
			} else {
				fmt.Println("Unexpected format for goals:", tds.Eq(5).Text())
			}
		}
	})

	clubStats := model.ClubStats{
		TeamName:    clubName,
		Season:      season,
		MarketValue: market_value,
		ResultStats: stats,
	}

	return clubStats, nil
}
