package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"freetchio-api/itch"
)

// getCategory serves all items of a itch.io category as JSON.
func getCategory(context *gin.Context, category itch.Category) {
	result, err := StorageService.Read(string(category))

	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	jsonData := []byte(result)
	context.Data(http.StatusOK, "application/json", jsonData)
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

// StartScrap scraps itch.io if the header is valid.
func StartScrap(context *gin.Context) {
	// Check if header is correct.
	cronKey := context.GetHeader("X-Cron-key")

	if cronKey != CRON_SCRAP_KEY {
		// Might as well do some ads :P
		msg := "Nope. " +
			"If you want to scrap more often, feel free to host it yourself. " +
			"You can find the code at https//github.com/shaigrorb/freetchio."
		context.String(http.StatusBadRequest, msg)
		return
	}

	// Sends accepted status as response since scrapping will take time.
	context.Status(http.StatusAccepted)

	// Update all items for each category.
	go scrapItchio(StorageService.Update)
}
