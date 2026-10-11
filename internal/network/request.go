package network

import (
	"context"
	"io"
	"net/http"
	"scraper/internal/logger"
	"time"
)

func MakeRequest(ctx context.Context, url string) ([]byte, error) {
	client := NewHTTPClient()

	var (
		attempt int
		resp    *http.Response
		req     *http.Request
		err     error
	)

	for {
		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)

		if err != nil {
			return nil, err
		}

		resp, err = client.Do(req)

		if err != nil {
			if attempt == maxAttempts {
				logger.Warnf("Can reach to site %s\n", url)
				return nil, err
			}

			if err = ctx.Err(); err != nil {
				return nil, err
			}

			attempt++
			logger.Warnf("Failed reach to site %s. Attempt %d/%d\n", url, attempt, maxAttempts)

			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(retryTimeout):
			}

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
