package controller

import (
	"net/http"
	"strconv"

	"github.com/franciscomp481/zerozero-stats-api/model"
	"github.com/franciscomp481/zerozero-stats-api/usecase"
	"github.com/gin-gonic/gin"
)

type stats_controller struct {
	usecase usecase.StatsUsecase
}

func NewStatsController(usecase usecase.StatsUsecase) stats_controller {
	return stats_controller{usecase: usecase}
}

func (s *stats_controller) GetPlayerStats(ctx *gin.Context) {

	playerName, index := ctx.Query("name"), ctx.Query("index")

	index_int, err := strconv.Atoi(index)
	if err != nil {
		response := model.Response{
			Message: "index must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return

	}

	if playerName == "" {
		response := model.Response{
			Message: "name parameter is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	playerStats, err := s.usecase.GetPlayerStats(playerName, index_int)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if playerStats.PlayerName == "" || len(playerStats.Seasons) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Player not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, playerStats)
}
