package itch

import (
	"encoding/json"
	"fmt"
)

type PageContent struct {
	Page     int    `json:"page"`
	NumItems int    `json:"num_items"`
	Content  string `json:"content"`
}

func (content *PageContent) Print() {
	fmt.Printf("Page: %d, Nb items: %d\n", content.Page, content.NumItems)
}

func (content *PageContent) FromJSON(j string) error {
	if err := json.Unmarshal([]byte(j), &content); err != nil {
		return err
	} else {
		return nil
	}
}
