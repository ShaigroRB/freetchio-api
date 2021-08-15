package main

import (
	"fmt"
	"freetchio-api/service"
	"os"

	"github.com/gin-gonic/gin"
)

// Environment variables
var (
	JSONBIN_API_KEY     = os.Getenv("JSONBIN_API_KEY")
	JSONBIN_INFO_BIN_ID = os.Getenv("JSONBIN_INFO_BIN_ID")
	CRON_SCRAP_KEY      = os.Getenv("CRON_SCRAP_KEY")
	PORT                = os.Getenv("PORT")
)

// Storage service used. Currently, it is jsonbin.io service.
var StorageService = service.JsonBin{
	ApiKey:    JSONBIN_API_KEY,
	InfoBinId: JSONBIN_INFO_BIN_ID,
	BinsIds:   &service.BinsIDs{},
}

// Storage service for local development
// var StorageService = service.FileService{}

func main() {
	// This is just because I'm too lazy to put all ids in environment variables.
	// And who knows, maybe future categories will appear.
	// Get all the bins ids thanks to the info bin id.
	err := StorageService.GetAllBinsIds()
	if err != nil {
		fmt.Println(err)
	}

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

	router.Run(":" + PORT)
}
