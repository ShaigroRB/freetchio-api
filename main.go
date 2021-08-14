package main

import "github.com/gin-gonic/gin"
// Environment variables
var (
	JSONBIN_API_KEY     = os.Getenv("JSONBIN_API_KEY")
	JSONBIN_INFO_BIN_ID = os.Getenv("JSONBIN_INFO_BIN_ID")
	CRON_SCRAP_KEY      = os.Getenv("CRON_SCRAP_KEY")
)

func main() {
	// coroutine that takes care of scrapping itch.io and creating the JSON files
	go ScrapItchioEvery12Hours()

	// Set the API.
	router := gin.Default()

	router.GET("/game-assets", GetGameAssets)
	router.GET("/books", GetBooks)
	router.GET("/comics", GetComics)
	router.GET("/tools", GetTools)
	router.GET("/games", GetGames)
	router.GET("/physical-games", GetPhysicalGames)
	router.GET("/soundtracks", GetSoundtracks)
	router.GET("/game-mods", GetGameMods)
	router.GET("/misc", GetMisc)

	router.POST("/scrap", StartScrap)

	router.Run()
}
