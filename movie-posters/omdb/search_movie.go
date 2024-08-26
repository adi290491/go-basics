package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchMovie(title string) (*Movie, error) {

	//apiKey := url.QueryEscape(omdb_api_key)
	//title = url.QueryEscape(title)
	url := OMDBURL + "?apikey=" + omdb_api_key + "&t=" + title
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var searchResult Movie

	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	return &searchResult, nil
}
