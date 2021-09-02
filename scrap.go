package main

import (
	"fmt"

	fditch "github.com/ShaigroRB/go-free-discount-itch"
)

// WriteItemsFn is a common type for Create & Update of the ItemService interface
type WriteItemsFn func(string, string) error

// writeValueForCategory writes value for a category
func writeValueForCategory(category fditch.Category, value string, fn WriteItemsFn) {
	key := string(category)
	err := fn(key, value)
	if err != nil {
		fmt.Println(err)
	}
}

// scrapItchio creates new JSON files for all the free on-sales items of itch.io.
func scrapItchio(fn WriteItemsFn) {
	for _, category := range fditch.Categories {
		jsonString := fditch.GetCategoryItemsAsJSON(category)
		writeValueForCategory(category, jsonString, fn)
	}
}
