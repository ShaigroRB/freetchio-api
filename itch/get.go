package itch

import (
	"fmt"
	"io"
	"net/http"
)

// itch.io hostname
const hostname = "https://itch.io"

// default parameters when doing API calls
const onsaleParams = "/on-sale?format=json&page"

// getPageJSON returns the content of a page for a given category.
// It returns the JSON as a string and an error if any.
func getPageJSON(category string, page int) (string, error) {
	url := fmt.Sprintf("%s/%s%s=%d", hostname, category, onsaleParams, page)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// getPageSales returns the content of a sales page and an error if any.
func getPageSales(link string) (string, error) {
	url := fmt.Sprintf("%s%s", hostname, link)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// getPageContent puts the content of a page in a channel for a given category.
// It returns whether it was the last page and an error if any.
func getPageContent(category string, page int, channel chan PageContent) (isLastPage bool, err error) {
	json, err := getPageJSON(category, page)
	if err != nil {
		fmt.Println(err)
		return isLastPage, err
	}

	content := PageContent{}
	err = content.FromJSON(json)
	if err != nil {
		fmt.Println(err)
		return isLastPage, err
	}
	channel <- content

	isLastPage = content.NumItems < 30

	return isLastPage, nil
}

// GetGameAssetsPageContent puts in a channel the `game-assets` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetGameAssetsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("game-assets", page, channel)
}

// GetBooksPageContent puts in a channel the `books` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetBooksPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("books", page, channel)
}

// GetToolsPageContent puts in a channel the `tools` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetToolsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("tools", page, channel)
}

// GetGamesPageContent puts in a channel the `games` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetGamesPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("games", page, channel)
}

// GetPhysicalGamesPageContent puts in a channel the `physical-games` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetPhysicalGamesPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("physical-games", page, channel)
}

// GetSoundstracksPageContent puts in a channel the `soundtracks` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetSoundtracksPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("soundtracks", page, channel)
}

// GetGameModsPageContent puts in a channel the `game-mods` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetGameModsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("game-mods", page, channel)
}

// GetMiscPageContent puts in a channel the `misc` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetMiscPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("misc", page, channel)
}
