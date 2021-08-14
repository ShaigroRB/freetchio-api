# Freetchio-api
This is the repository for the API of Freetchio.  
It is composed of:
- A scrapper package that exposes methods to get all the free items for all the categories.
- The API that exposes endpoints to get a JSON file containing information of free items for each category.

## Scrapper package
### TODO list
1. [ ] Get max amount of results
    1. [ ] Get items/on-sale
    2. [ ] Parse to get the number of results (`curl "https://itch.io/items/on-sale" | grep -i "<nobr class=\"game\_count\".*</nobr>"`)
2. [x] Get the json of all pages
    1. [x] Get category/on-sale?format=json&page=.. for each json
    2. [x] Put each one in a struct (PageContent)
3. [x] Parse the content as incomplete items (*missing the end date for sales*)
    1. [x] Read the content as html nodes
    2. [x] Construct items based on the nodes
        1. Split the nodes to keep only the "*game_cells*" nodes (*Spoiler: it's not worth it*)
        2. [x] Create items from those "*game_cells*" nodes. Don't forget to only keep the **100%** on sales items. (*careful of +100% sales*)
4. [x] Get the end date for each item
    1. [x] Get html content from the sales link
    2. [x] Parse it to get the end date for the sales (`{"start_date":"2021-05-28T10:00:35Z","id":50563,"end_date":"2021-05-30T10:02:59Z","can_be_bought":true,"actual_price":398}`)
5. [x] Create a JSON string out of all the items
6. [x] Keep it simple by removing all concurrency. Any concurrency should be done in the API.

### Structure of one cell
How I represent it:  
html_element interesting_attribute (why) [text] (??)  

- div data-game_id
    - a href (game link)
        - div data-background_image (image for the cell)
            - div ??
            - div data-gif (gif for the cell) ??
    - div
        - a
            - span
    - div
        - div
            - a href (game link) [game title]
            - a href (sales link)
                - div [price value]
                - div [sale tag]
        - div title (game description) [game description]
        - div
            - a href (game author) [game author]
        - div ?
            - span title (windows) ??
            - span title (linux) ??
            - span title (mac) ??


## API
The API only exposes GET endpoints:
- `/game-assets`    returns game assets
- `/books`          returns books
- `/comics`         returns comics
- `/tools`          returns tools
- `/games`          returns games
- `/physical-games` returns physical games
- `/soundtracks`    returns soundtracks
- `/game-mods`      returns game mods
- `/misc`           returns misc items

Each of those items are JSON objects of the form:
```json
{
    "id": "404609",
    "link": "https://dragonruby.itch.io/dragonruby-gtk",
    "img_link": "https://img.itch.zone/aW1nLzIzNjU2MzQucG5n/315x250%23c/Nrmttz.png",
    "title": "DragonRuby Game Toolkit",
    "description": "",
    "author": "DragonRuby",
    "sales_link": "/s/54971/lowrezjam-participation-sale",
    "end_date": "2021-08-07T05:00:00Z"
}
```


## License
This project is under the MIT license.