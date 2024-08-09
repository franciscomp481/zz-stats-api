package controller

import (
	"net/http"
	"strconv"

	"github.com/franciscomp481/zerozero-stats-api/model"
	"github.com/franciscomp481/zerozero-stats-api/usecase"
	"github.com/franciscomp481/zerozero-stats-api/webscrapper"
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

	fullURL, err := webscrapper.SearchPlayer(playerName, index_int)
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

	// Marshal playerStats to JSON
	//jsonData, err := json.MarshalIndent(playerStats, "", "  ")
	//if err != nil {
	//	log.Fatalf("Error marshalling playerStats to JSON: %v", err)
	//}

	ctx.JSON(http.StatusOK, playerStats)
}
