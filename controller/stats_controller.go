package controller

import "github.com/franciscomp481/zerozero-stats-api/usecase"

type StatsController struct {
	usecase usecase.StatsUsecase
}

func NewStatsController(usecase usecase.StatsUsecase) StatsController {
	return StatsController{usecase: usecase}
}
