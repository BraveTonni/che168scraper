package network

import (
	"encoding/json"
	"io"
	"net/http"
	"scraper/internal/model"
)

func MakeRequest(url string) (*model.SearchResponse, error) {
	client := NewHttpClient()
	resp, err := client.Get(url)

	var result model.SearchResponse

	if err != nil {
		return &result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &result, ErrFailReachSite
	}

	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &result); err != nil {
		return &result, err
	}

	return &result, nil
}
