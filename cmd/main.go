package main

import (
	"github.com/franciscomp481/zerozero-stats-api/controller"
	"github.com/franciscomp481/zerozero-stats-api/repository"
	"github.com/franciscomp481/zerozero-stats-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Repository initialization
	statsRepository := repository.NewStatsRepository()
	statsUsecase := usecase.NewStatsUsecase(statsRepository)
	statsController := controller.NewStatsController(statsUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/playerstats", statsController.GetPlayerStats)

	server.Run(":8080")
}
