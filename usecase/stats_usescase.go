package usecase

import "github.com/franciscomp481/zerozero-stats-api/repository"

type StatsUsecase struct {
	repository repository.StatsRepository
}

func NewStatsUsecase(repo repository.StatsRepository) StatsUsecase {
	return StatsUsecase{repository: repo}
}
