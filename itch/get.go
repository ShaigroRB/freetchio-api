package itch

import (
	"fmt"
	"io"
	"net/http"
)

const hostname = "https://itch.io/"
const onsaleParams = "/on-sale?format=json&page"

func getPageJSON(category string, page int) (string, error) {
	url := fmt.Sprintf("%s%s%s=%d", hostname, category, onsaleParams, page)
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

func GetGameAssetsPageContent(page int, channel chan PageContent) (isLastPage bool, err error) {
	json, err := getPageJSON("game-assets", page)
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

	if isLastPage {
		close(channel)
	}
	return isLastPage, nil
}
