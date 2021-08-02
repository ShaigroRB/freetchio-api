package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"freetchio-api/itch"
)

// writeJSONFile writes content into a JSON file.
func writeJSONFile(content string, filename string) error {
	jsonFilename := fmt.Sprintf("%s.json", filename)
	contentBytes := []byte(content)
	err := ioutil.WriteFile(jsonFilename, contentBytes, 0644)
	return err
}

// scrapItchio creates new JSON files for all the free on-sales items of itch.io.
func scrapItchio() {
	for _, category := range itch.Categories {
		jsonString := itch.GetCategoryItemsAsJSON(category)
		err := writeJSONFile(jsonString, string(category))

		if err != nil {
			fmt.Println(err)
		}
	}
}

// ScrapItchioEvery12Hours creates new JSON files every 12 hours for all the free on-sale items of itch.io.
func ScrapItchioEvery12Hours() {
	// Create the JSON files for the first time
	scrapItchio()

	for range time.Tick(time.Hour * 12) {
		scrapItchio()
	}
}
