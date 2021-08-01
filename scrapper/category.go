package scrapper

type Category int

// Enum for all categories that can be found on itch.io
const (
	GameAssets Category = iota
	Books
	Tools
	Games
	PhysicalGames
	Soundtracks
	GameMods
	Misc
)
