package main

import (
	"github.com/franciscomp481/zerozero-stats-api/controller"
	_ "github.com/franciscomp481/zerozero-stats-api/docs" // Import the generated docs package
	"github.com/franciscomp481/zerozero-stats-api/repository"
	"github.com/franciscomp481/zerozero-stats-api/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("/playerstats", statsController.GetPlayerStats)

	server.Run(":8080")
}
