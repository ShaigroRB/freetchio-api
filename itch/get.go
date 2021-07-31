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

// getPageContent puts the content of a page in a list for a given category.
// It returns whether it was the last page and an error if any.
func getPageContent(category string, page int, list *[]PageContent) (isLastPage bool, err error) {
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

	*list = append(*list, content)

	isLastPage = content.NumItems < 30

	return isLastPage, nil
}

// Type that represents a function to get a PageContent for a specific category.
type GetCategoryPageContentFn func(int, *[]PageContent) (bool, error)

// GetGameAssetsPageContent puts in a list the `game-assets` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetGameAssetsPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("game-assets", page, list)
}

// GetBooksPageContent puts in a list the `books` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetBooksPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("books", page, list)
}

// GetToolsPageContent puts in a list the `tools` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetToolsPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("tools", page, list)
}

// GetGamesPageContent puts in a list the `games` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetGamesPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("games", page, list)
}

// GetPhysicalGamesPageContent puts in a list the `physical-games` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetPhysicalGamesPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("physical-games", page, list)
}

// GetSoundstracksPageContent puts in a list the `soundtracks` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetSoundtracksPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("soundtracks", page, list)
}

// GetGameModsPageContent puts in a list the `game-mods` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetGameModsPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("game-mods", page, list)
}

// GetMiscPageContent puts in a list the `misc` type content for a given page.
// It returns whether it was the last pageand an error if any.
func GetMiscPageContent(page int, list *[]PageContent) (isLastPage bool, err error) {
	return getPageContent("misc", page, list)
}
