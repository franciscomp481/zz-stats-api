package main

import (
	"encoding/json"
	"fmt"

	"github.com/franciscomp481/zerozero-stats-api/model"
	"github.com/franciscomp481/zerozero-stats-api/webscrapper"
)

func main() {
	club := model.ClubFilters{
		ClubName: "Sporting",
		Index:    0,
	}

	clubURL, err := webscrapper.SearchClub(club)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := webscrapper.GetPage(clubURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	clubStats, err := webscrapper.FetchClubStats(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert clubStats to JSON
	clubStatsJSON, err := json.MarshalIndent(clubStats, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling clubStats to JSON:", err)
		return
	}

	// Print the JSON string
	fmt.Println(string(clubStatsJSON))

}
