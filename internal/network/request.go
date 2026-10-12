package network

import (
	"context"
	"fmt"
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
				return nil, fmt.Errorf("make request: %w", ErrFailReachSite)
			}

			if err = ctx.Err(); err != nil {
				return nil, fmt.Errorf("make request: %w", ErrFailReachSite)
			}

			attempt++
			logger.Warnf("Failed reach to site %s. Attempt %d/%d\n", url, attempt, maxAttempts)

			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("make request: %w", ErrFailReachSite)
			case <-time.After(retryTimeout):
			}

			continue
		}

		break
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("make request, status: %d: %w", resp.StatusCode, ErrFailReachSite)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	return body, err
}
