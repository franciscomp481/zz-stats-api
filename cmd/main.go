package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/franciscomp481/zerozero-stats-api/webscrapper"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Repository initialization
	//statsRepository := repository.NewStatsRepository()
	//statsUsecase := usecase.NewStatsUsecase(statsRepository)
	//statsController := controller.NewStatsController(statsUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	playerName, index := "Pepe", 0
	fullURL, err := webscrapper.SearchPlayer(playerName, index)
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
	jsonData, err := json.MarshalIndent(playerStats, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling playerStats to JSON: %v", err)
	}

	fmt.Println(string(jsonData))

	//server.Run(":8081")
}
