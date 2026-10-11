package network

import (
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

func NewHTTPClient() *http.Client {
	return &http.Client{Timeout: defaultTimeout}
}
