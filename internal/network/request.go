package network

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"scraper/internal/logger"
	"time"
)

func MakeRequest(ctx context.Context, url string) ([]byte, error) {

	resp, err := requestWithRetry(ctx, url)

	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("make request, status: %d: %w", resp.StatusCode, ErrFailReachSite)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	return body, nil
}

func requestWithRetry(ctx context.Context, url string) (*http.Response, error) {
	client := NewHTTPClient()

	var lastErr error // store last error if all retries failed

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

		if err != nil {
			return nil, fmt.Errorf("retry: %w", err)
		}

		resp, err := client.Do(req)

		// successful reach to site
		if err == nil {
			return resp, nil
		}
		lastErr = err

		if attempt == maxAttempts {
			return nil, fmt.Errorf("retry after %d attempts: %w",
				maxAttempts,
				errors.Join(lastErr, ErrFailReachSite),
			)
		}

		logger.Debugf("Failed reach to site %s. Attempt %d/%d\n", url, attempt, maxAttempts)

		// retry timeout
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryTimeout):
		}
	}

	return nil, fmt.Errorf("retry after %d attempts: %w",
		maxAttempts,
		errors.Join(lastErr, ErrFailReachSite),
	)
}
