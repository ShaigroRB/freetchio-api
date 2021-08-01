package scrapper

type Category string

// Enum for all categories that can be found on itch.io
const (
	GameAssets    Category = "game-assets"
	Books                  = "books"
	Tools                  = "tools"
	Games                  = "games"
	PhysicalGames          = "physical-games"
	Soundtracks            = "soundstracks"
	GameMods               = "game-mods"
	Misc                   = "misc"
)

// Array containing all categories.
var Categories = []Category{
	GameAssets,
	Books,
	Tools,
	Games,
	PhysicalGames,
	Soundtracks,
	GameMods,
	Misc,
}
