package repository

import (
	"github.com/franciscomp481/zerozero-stats-api/model"
	"github.com/franciscomp481/zerozero-stats-api/webscrapper"
)

type StatsRepository struct {
	// insert webscrapper here
}

func NewStatsRepository() StatsRepository {
	return StatsRepository{}
}

func (r *StatsRepository) GetPlayerStats(playerName string, index int) (*model.PlayerStats, error) {
	// insert webscrapper here

	fullURL, err := webscrapper.SearchPlayer(playerName, index)
	if err != nil {
		panic(err)
	}

	doc, err := webscrapper.GetPlayerPage(fullURL)
	if err != nil {
		panic(err)
	}

	playerStats, err := webscrapper.FetchPlayerStats(doc)

	if err != nil {
		panic(err)
	}

	return &playerStats, nil
}
