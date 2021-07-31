package itch

import (
	"fmt"
	"io"
	"net/http"
)

const hostname = "https://itch.io"
const onsaleParams = "/on-sale?format=json&page"

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

func GetGameAssetsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("game-assets", page, channel)
}

func GetBooksPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("books", page, channel)
}

func GetToolsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("tools", page, channel)
}

func GetGamesPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("games", page, channel)
}

func GetPhysicalGamesPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("physical-games", page, channel)
}

func GetSoundtracksPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("soundtracks", page, channel)
}

func GetGameModsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("game-mods", page, channel)
}

func GetMiscPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	return getPageContent("misc", page, channel)
}
