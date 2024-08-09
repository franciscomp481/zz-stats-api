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

func (u *StatsUsecase) GetPlayerStats(playerName string, index int) (*model.PlayerStats, error) {
	playerStats, err := u.repository.GetPlayerStats(playerName, index)

	if err != nil {
		return nil, err
	}

	return playerStats, nil
}
