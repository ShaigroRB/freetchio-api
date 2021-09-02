# Freetchio-api
This is the repository for the API of Freetchio.  
It exposes endpoints to get a JSON file containing information of free items for each category.

## Scrapper
The scrapper uses the `fditch` package from [this package](https://github.com/ShaigroRB/go-free-discount-itch).

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