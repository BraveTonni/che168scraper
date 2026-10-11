package network

import (
	"net/http"
)

func NewHTTPClient() *http.Client {
	return &http.Client{Timeout: defaultTimeout}
}
