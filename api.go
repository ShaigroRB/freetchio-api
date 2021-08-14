package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"freetchio-api/itch"
)

// getCategory serves all items of a itch.io category as JSON.
func getCategory(context *gin.Context, category itch.Category) {
	filename := fmt.Sprintf("%s.json", category)
	context.File(filename)
}

// GetGameAssets serves all game assets items as JSON.
func GetGameAssets(context *gin.Context) {
	getCategory(context, itch.GameAssets)
}

// GetBooks serves all books items as JSON.
func GetBooks(context *gin.Context) {
	getCategory(context, itch.Books)
}

// GetComics serves all comics items as JSON.
func GetComics(context *gin.Context) {
	getCategory(context, itch.Comics)
}

// GetTools serves all tools items as JSON.
func GetTools(context *gin.Context) {
	getCategory(context, itch.Tools)
}

// GetGames serves all games items as JSON.
func GetGames(context *gin.Context) {
	getCategory(context, itch.Games)
}

// GetPhysicalGames serves all physical games items as JSON.
func GetPhysicalGames(context *gin.Context) {
	getCategory(context, itch.PhysicalGames)
}

// GetSoundtracks serves all soundtracks items as JSON.
func GetSoundtracks(context *gin.Context) {
	getCategory(context, itch.Soundtracks)
}

// GetGameMods serves all game mods items as JSON.
func GetGameMods(context *gin.Context) {
	getCategory(context, itch.GameMods)
}

// GetMisc serves all misc items as JSON.
func GetMisc(context *gin.Context) {
	getCategory(context, itch.Misc)
}
