package itch

import (
	"encoding/json"
	"fmt"
)

// arrayToJSON converts an array to a JSON string.
func arrayToJSON(items []Item) (string, error) {
	result, err := json.Marshal(items)
	return string(result), err
}

// getCategoryAllPageContents returns a channel containing all the PageContent for a category.
// It returns an error with it if any.
func getCategoryAllPageContents(getCategoryPageContentFn GetCategoryPageContentFn) (chan PageContent, error) {
	var err error

	channel := make(chan PageContent)
	go func() {
		defer close(channel)

		page := 0
		for {
			page++
			isLastPage, err := getCategoryPageContentFn(page, channel)
			if err != nil || isLastPage {
				break
			}
		}
	}()

	return channel, err
}

func errorToJSON(err error) string {
	return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
}

// getCategoryItemsAsJSON returns a JSON string containing all items using a function to get the page contents for a category.
func getCategoryItemsAsJSON(getCategoryPageContentFn GetCategoryPageContentFn) string {
	// get a channel with all the PageContent for a category
	pageContentChannel, err := getCategoryAllPageContents(getCategoryPageContentFn)

	if err != nil {
		return errorToJSON(err)
	}

	var items []Item

	for {
		// loop on the element in the channel
		content, ok := <-pageContentChannel
		if !ok {
			break
		}

		// for each PageContent, parse the items in it
		itemsForPageContent, err := ConvertPageContentToItems(content)

		if err != nil {
			return errorToJSON(err)
		}

		// add those items to the list of all the items
		for item := range itemsForPageContent {
			items = append(items, item)
		}
	}

	// convert all the items to a single json
	result, err := arrayToJSON(items)

	if err != nil {
		return errorToJSON(err)
	}
	return result
}

// GetCategoryItemsAsJSON returns a JSON string containing all items for a given category.
func GetCategoryItemsAsJSON(category Category) string {
	var itemsAsJSON string

	// inner function to keep readability in the next lines
	setResult := func(fn GetCategoryPageContentFn) {
		itemsAsJSON = getCategoryItemsAsJSON(fn)
	}

	switch category {
	case GameAssets:
		setResult(GetGameAssetsPageContent)
	case Books:
		setResult(GetBooksPageContent)
	case Tools:
		setResult(GetToolsPageContent)
	case Games:
		setResult(GetGamesPageContent)
	case PhysicalGames:
		setResult(GetPhysicalGamesPageContent)
	case Soundtracks:
		setResult(GetSoundtracksPageContent)
	case GameMods:
		setResult(GetGameModsPageContent)
	case Misc:
		setResult(GetMiscPageContent)
	}

	return itemsAsJSON
}
