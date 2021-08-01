package main

import "github.com/gin-gonic/gin"

func main() {
	// coroutine that takes care of scrapping itch.io and creating the JSON files
	go ScrapItchioEvery12Hours()

	// set the API
	router := gin.Default()

	router.GET("/game-assets", GetGameAssets)
	router.GET("/books", GetBooks)
	router.GET("/tools", GetTools)
	router.GET("/games", GetGames)
	router.GET("/physical-games", GetPhysicalGames)
	router.GET("/soundtracks", GetSoundtracks)
	router.GET("/game-mods", GetGameMods)
	router.GET("/misc", GetMisc)

	router.Run("localhost:8042")
}
