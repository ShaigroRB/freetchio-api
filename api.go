package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"freetchio-api/scrapper"
)

// getCategory serves all items of a itch.io category as JSON.
func getCategory(context *gin.Context, category scrapper.Category) {
	filename := fmt.Sprintf("%s.json", category)
	context.File(filename)
}

// GetGameAssets serves all game assets items as JSON.
func GetGameAssets(context *gin.Context) {
	getCategory(context, scrapper.GameAssets)
}

// GetBooks serves all books items as JSON.
func GetBooks(context *gin.Context) {
	getCategory(context, scrapper.Books)
}

// GetTools serves all tools items as JSON.
func GetTools(context *gin.Context) {
	getCategory(context, scrapper.Tools)
}

// GetGames serves all games items as JSON.
func GetGames(context *gin.Context) {
	getCategory(context, scrapper.Games)
}

// GetPhysicalGames serves all physical games items as JSON.
func GetPhysicalGames(context *gin.Context) {
	getCategory(context, scrapper.PhysicalGames)
}

// GetSoundtracks serves all soundtracks items as JSON.
func GetSoundtracks(context *gin.Context) {
	getCategory(context, scrapper.Soundtracks)
}

// GetGameMods serves all game mods items as JSON.
func GetGameMods(context *gin.Context) {
	getCategory(context, scrapper.GameMods)
}

// GetMisc serves all misc items as JSON.
func GetMisc(context *gin.Context) {
	getCategory(context, scrapper.Misc)
}
