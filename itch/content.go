package itch

import (
	"encoding/json"
	"fmt"
)

// API calls to itch.io return a JSON object.
// The PageContent struct is based on that JSON object.
type PageContent struct {
	Page     int    `json:"page"`
	NumItems int    `json:"num_items"`
	Content  string `json:"content"`
}

// Print prints the PageContent struct.
func (content *PageContent) Print() {
	fmt.Printf("Page: %d, Nb items: %d\n", content.Page, content.NumItems)
}

// FromJSON deserializes a JSON and puts it into the PageContent struct.
func (content *PageContent) FromJSON(j string) error {
	if err := json.Unmarshal([]byte(j), &content); err != nil {
		return err
	} else {
		return nil
	}
}
