# Freetchio-api
This is the repository for the API of Freetchio.  
It exposes endpoints to get a JSON file containing information of free items for each category.

## Scrapper
The scrapper uses the `fditch` package from [this package](https://github.com/ShaigroRB/go-free-discount-itch).

## Endpoints

There are two types of endpoints.  

1. The GET endpoints return items for a specific category:
    - `/game-assets`    returns game assets
    - `/books`          returns books
    - `/comics`         returns comics
    - `/tools`          returns tools
    - `/games`          returns games
    - `/physical-games` returns physical games
    - `/soundtracks`    returns soundtracks
    - `/game-mods`      returns game mods
    - `/misc`           returns misc items
2. The POST endpoint `/scrap` runs the scrapper that will get all items for all categories and use an *ItemService* to write the result somewhere.

## Items
Items returned by the GET endpoints are of the following format:
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

## How to use the API locally

To use the API, you can either:
1. run it using local files
2. run it using jsonbin.io

### 1. Run it using local files
#### Environment variables
- **PORT**: The port on which the API is exposed.
- **CRON_SCRAP_KEY**: The key used to avoid that everyone runs the scrapper whenever they want.

#### Update main.go
Comment and uncomment a few things in `main.go`:
```go
...
// Storage service used. Currently, it is jsonbin.io service.
// var StorageService = service.JsonBin{
// 	ApiKey:    JSONBIN_API_KEY,
// 	InfoBinId: JSONBIN_INFO_BIN_ID,
// 	BinsIds:   &service.BinsIDs{},
// }

// Storage service for local development
var StorageService = service.FileService{}

func main() {
	// This is just because I'm too lazy to put all ids in environment variables.
	// And who knows, maybe future categories will appear.
	// Get all the bins ids thanks to the info bin id.
	// err := StorageService.GetAllBinsIds()
	// if err != nil {
	// 	fmt.Println(err)
    // }

    // Set the API.
	router := gin.Default()
    ...
}
```

#### Start the API
- Set the environment variables:
    ```bash
    export PORT=8080;
    export CRON_SCRAP_KEY="notASecure1"
    ```
- Start the API: `go run main.go scrap.go service.go api.go`
- Then run the scrapper once: `curl -v -H "X-Cron-key: notASecure1" --request POST localhost:8080/scrap`
- Then enjoy hitting on the endpoints: `curl -v --request GET localhost:8080/game-assets`

### 2. Run it using jsonbin.io

#### Environment variables
- **PORT**: The port on which the API is exposed.
- **CRON_SCRAP_KEY**: The key used to avoid that everyone runs the scrapper whenever they want.
- **JSONBIN_API_KEY**: The API key to write the data at jsonbin.io
- **JSONBIN_INFO_BIN_ID**: The bin ID that contains all the bin IDs for each category (*this is because I'm too lazy to write each bin ID as an environment variable*)

#### Start the API
- Set the environment variables:
    ```bash
    export PORT=8080;
    export CRON_SCRAP_KEY="notASecure1";
    export JSONBIN_API_KEY="justARandomKey";
    export JSONBIN_INFO_BIN_ID="andAnother1"
    ```
- Start the API: `go run main.go scrap.go service.go api.go`
- Then run the scrapper once: `curl -v -H "X-Cron-key: notASecure1" --request POST localhost:8080/scrap`
- Then enjoy hitting on the endpoints: `curl -v --request GET localhost:8080/game-assets`

## License
This project is under the MIT license.