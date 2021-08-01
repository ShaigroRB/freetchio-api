package main

import (
	"fmt"
	"freetchio-api/scrapper"
)

func main() {
	fmt.Println(scrapper.GetCategoryItemsAsJSON(scrapper.Tools))
}
