package controller

import (
	"net/http"
	"strconv"

	"github.com/franciscomp481/zerozero-stats-api/model"
	"github.com/franciscomp481/zerozero-stats-api/usecase"
	"github.com/gin-gonic/gin"
)

// stats_controller handles the stats usecase
type stats_controller struct {
	usecase usecase.StatsUsecase
}

// NewStatsController creates a new stats controller
func NewStatsController(usecase usecase.StatsUsecase) stats_controller {
	return stats_controller{usecase: usecase}
}

// GetPlayerStats godoc
// @Summary Get player statistics
// @Description Get statistics for a player based on name, index, and nationality
// @Tags stats
// @Accept  json
// @Produce  json
// @Param name query string true "Player Name"
// @Param index query string true "Index"
// @Param nationality query string false "Nationality"
// @Success 200 {object} model.PlayerStats
// @Failure 400 {object} model.Response
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /playerstats [get]
func (s *stats_controller) GetPlayerStats(ctx *gin.Context) {

	playerName, index, nationality := ctx.Query("name"), ctx.Query("index"), ctx.Query("nationality")

	if index == "" {
		index = "0"
	}

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

	playerFilters := model.PlayerFilters{
		PlayerName:  playerName,
		Index:       index_int,
		Nationality: nationality,
	}

	playerStats, err := s.usecase.GetPlayerStats(playerFilters)

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
