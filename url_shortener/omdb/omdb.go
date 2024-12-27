package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Search       []SearchResult
	TotalResults string
	Response     string
}

type SearchResult struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

func Search(apiKey, title string) (Result, error) {
	var v url.Values

	v.Set("apikey", apiKey)
	v.Set("s", title)

	resp, err := http.Get("http://www.obmdapi.com/" + v.Encode())

	if err != nil {
		return Result{}, fmt.Errorf("failed to make request to omdb: %w", err)
	}
	defer resp.Body.Close()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}