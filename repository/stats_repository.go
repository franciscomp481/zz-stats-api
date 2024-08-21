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

func (r *StatsRepository) GetPlayerStats(filters model.PlayerFilters) (*model.PlayerStats, error) {
	// insert webscrapper here

	fullURL, err := webscrapper.SearchPlayer(filters)
	if err != nil {
		panic(err)
	}

	doc, err := webscrapper.GetPage(fullURL)
	if err != nil {
		panic(err)
	}

	playerStats, err := webscrapper.FetchPlayerStats(doc)

	if err != nil {
		panic(err)
	}

	return &playerStats, nil
}

func (r *StatsRepository) GetClubStats(filters model.ClubFilters) (*model.ClubStats, error) {

	fullURL, err := webscrapper.SearchClub(filters)
	if err != nil {
		panic(err)
	}

	doc, err := webscrapper.GetPage(fullURL)
	if err != nil {
		panic(err)
	}

	clubStats, err := webscrapper.FetchClubStats(doc)

	if err != nil {
		panic(err)
	}

	return &clubStats, nil
}
