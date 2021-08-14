package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// jsonbin.io default endpoint
const jsonbinUrl = "https://api.jsonbin.io/v3/b"

// BinsIDs contains all the bin ids for categories
type BinsIDs struct {
	GameAssets    string `json:"game_assets"`
	Books         string `json:"books"`
	Comics        string `json:"comics"`
	Tools         string `json:"tools"`
	Games         string `json:"games"`
	PhysicalGames string `json:"physical_games"`
	Soundtracks   string `json:"soundtracks"`
	GameMods      string `json:"game_mods"`
	Misc          string `json:"misc"`
}

// A JsonBin contains all the info needed to send data to jsonbin.io.
type JsonBin struct {
	ApiKey    string
	InfoBinId string
	BinsIds   *BinsIDs
}

// getKeyID returns the bin ID corresponding to a given key.
func (bin *JsonBin) getBinID(key string) string {
	var id string

	binsIds := bin.BinsIds

	switch key {
	case "game-assets":
		id = binsIds.GameAssets
	case "books":
		id = binsIds.Books
	case "comics":
		id = binsIds.Comics
	case "tools":
		id = binsIds.Tools
	case "games":
		id = binsIds.Games
	case "physical-games":
		id = binsIds.PhysicalGames
	case "soundtracks":
		id = binsIds.Soundtracks
	case "game-mods":
		id = binsIds.GameMods
	case "misc":
		id = binsIds.Misc
	default:
		id = key
	}

	return id
}

// requestWithApiKey does a request with an api key as header.
func requestWithApiKey(req *http.Request, apiKey string) error {
	req.Header.Add("X-Master-key", apiKey)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil
}

// requestWithApiKeyWithResponse does a request with an api key as header and returns the reponse.
func requestWithApiKeyWithResponse(req *http.Request, apiKey string) (*http.Response, error) {
	req.Header.Add("X-Master-key", apiKey)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return resp, nil
}

// GetAllBinsIds gets all the id of each bin.
func (bin *JsonBin) GetAllBinsIds() error {
	infoBin, err := bin.Read(bin.InfoBinId)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(infoBin), &bin.BinsIds)

	if err != nil {
		return err
	} else {
		return nil
	}
}

// Create creates a jsonbin given a key & a value.
func (bin *JsonBin) Create(key string, value string) error {
	apiKey := bin.ApiKey
	valueBytes := bytes.NewBufferString(value)

	req, err := http.NewRequest("POST", jsonbinUrl, valueBytes)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	return requestWithApiKey(req, apiKey)
}

// Update updates a jsonbin thanks to a key.
func (bin *JsonBin) Update(key string, value string) error {
	apiKey := bin.ApiKey
	valueBytes := bytes.NewBufferString(value)
	binID := bin.getBinID(key)
	url := fmt.Sprintf("%s/%s", jsonbinUrl, binID)

	req, err := http.NewRequest("PUT", url, valueBytes)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	return requestWithApiKey(req, apiKey)
}

// Read reads a jsonbin thanks to a key.
func (bin *JsonBin) Read(key string) (string, error) {
	apiKey := bin.ApiKey
	binID := bin.getBinID(key)
	url := fmt.Sprintf("%s/%s/latest", jsonbinUrl, binID)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	resp, err := requestWithApiKeyWithResponse(req, apiKey)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
