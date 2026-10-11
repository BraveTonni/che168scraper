package network

import (
	"io"
	"net/http"
	"scraper/internal/logger"
	"time"
)

func MakeRequest(url string) ([]byte, error) {
	client := NewHTTPClient()

	var (
		attempt int
		resp    *http.Response
		err     error
	)

	for {
		resp, err = client.Get(url)

		if err != nil {
			if attempt == maxAttempts {
				logger.Warnf("Can reach to site %s\n", url)
				return nil, err
			}

			attempt++
			logger.Warnf("Failed reach to site %s. Attempt %d/%d\n", url, attempt, maxAttempts)
			time.Sleep(retryTimeout)
			continue
		}

		break
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrFailReachSite
	}

	body, err := io.ReadAll(resp.Body)

	return body, err
}
