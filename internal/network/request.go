package network

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
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
		resp, err := doOnce(ctx, client, url)

		if err == nil {
			return resp, nil
		}
		lastErr = err

		if attempt < maxAttempts {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(retryTimeout):
			}
		}
	}

	return nil, fmt.Errorf("retry after %d attempts: %w",
		maxAttempts,
		errors.Join(lastErr, ErrFailReachSite),
	)
}

func doOnce(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	return resp, nil
}
