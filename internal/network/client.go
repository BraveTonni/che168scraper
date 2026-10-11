package network

import (
	"net/http"
	"sync"
	"time"
)

var (
	httpClient *http.Client
	once       sync.Once
)

func NewHttpClient() *http.Client {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	})
	return httpClient
}
