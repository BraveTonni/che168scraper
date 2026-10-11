package network

import (
	"io"
	"net/http"
)

func MakeRequest(url string) ([]byte, error) {
	client := NewHTTPClient()
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrFailReachSite
	}

	body, err := io.ReadAll(resp.Body)

	return body, err
}
