package itch

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Item struct {
	ID          string
	Link        string
	ImgLink     string
	Title       string
	Description string
	Author      string
	SalesLink   string
	EndDate     string
}

func getPreOrderQueue(root *html.Node) []*html.Node {
	queue := make([]*html.Node, 0)
	res := make([]*html.Node, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		res = append(res, node)

		for child := node.LastChild; child != nil; child = child.PrevSibling {
			queue = append(queue, child)
		}
	}

	return res
}

func nodeToItemsWithoutEndDate(root *html.Node, maxItems int) chan Item {
	cells := make(chan Item, maxItems)
	nodes := getPreOrderQueue(root)

	var cell Item

	isNodeAGame := func(n *html.Node) bool {
		if len(n.Attr) == 0 {
			return false
		}
		return n.DataAtom == atom.Div && n.Attr[0].Key == "data-game_id"
	}

	continueTillNextGame := true

	for _, node := range nodes {
		if continueTillNextGame && !isNodeAGame(node) {
			continue
		} else {
			continueTillNextGame = false
		}

		switch node.DataAtom {
		case atom.Div:
			if len(node.Attr) > 0 {
				if attr := node.Attr[0]; attr.Key == "data-game_id" {
					if cell.ID != "" {
						cells <- cell
					}
					cell = Item{}
					cell.ID = attr.Val
				} else if attr.Key == "class" && attr.Val == "game_author" {
					cell.Author = node.FirstChild.FirstChild.Data
				} else if attr.Key == "class" &&
					(attr.Val == "sale_tag" || attr.Val == "sale_tag reverse_sale") {
					// cause yes, reverse sales are a thing in itch.io
					if node.FirstChild.Data != "-100%" {
						continueTillNextGame = true
						cell.ID = ""
					}
				} else if len(node.Attr) > 1 && node.Attr[1].Key == "data-background_image" {
					cell.ImgLink = node.Attr[1].Val
				}
			}
		case atom.A:
			if len(node.Attr) > 1 {
				if attr := node.Attr[0]; attr.Key == "class" && attr.Val == "title game_link" {
					cell.Link = node.Attr[1].Val
					cell.Title = node.FirstChild.Data
				} else if len(node.Attr) > 2 {
					if attr = node.Attr[2]; attr.Key == "class" && attr.Val == "price_tag meta_tag sale" {
						cell.SalesLink = node.Attr[0].Val
					}
				}
			}
		default:
			continue
		}
	}

	close(cells)
	return cells
}

func parseEndDate(body string) string {
	regx := regexp.MustCompile(`end_date\".*\",`)
	matches := regx.FindStringSubmatch(body)
	regx = regexp.MustCompile(`[0-9]+-[^\"]*`)
	matches = regx.FindStringSubmatch(matches[0])

	return matches[0]
}

func PageContentToItems(content PageContent) (chan Item, error) {
	reader := strings.NewReader(content.Content)
	node, err := html.Parse(reader)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	partialItems := nodeToItemsWithoutEndDate(node, content.NumItems)
	items := make(chan Item, len(partialItems))

	defer close(items)

	for partialItem := range partialItems {
		body, err := getPageSales(partialItem.SalesLink)
		if err != nil {
			fmt.Print(err)
			return items, err
		}

		partialItem.EndDate = parseEndDate(body)

		items <- partialItem
	}

	return items, nil
}