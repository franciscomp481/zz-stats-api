package usecase

import (
	"github.com/franciscomp481/zerozero-stats-api/model"
	"github.com/franciscomp481/zerozero-stats-api/repository"
)

type StatsUsecase struct {
	repository repository.StatsRepository
}

func NewStatsUsecase(repo repository.StatsRepository) StatsUsecase {
	return StatsUsecase{repository: repo}
}

func (u *StatsUsecase) GetPlayerStats(filters model.PlayerFilters) (*model.PlayerStats, error) {
	playerStats, err := u.repository.GetPlayerStats(filters)

	if err != nil {
		return nil, err
	}

	return playerStats, nil
}

func (u *StatsUsecase) GetClubStats(filters model.ClubFilters) (*model.ClubStats, error) {
	clubStats, err := u.repository.GetClubStats(filters)

	if err != nil {
		return nil, err
	}

	return clubStats, nil
}
