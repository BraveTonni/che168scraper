package scraper

import (
	"context"
	"fmt"
	"scraper/internal/dto"
	"scraper/internal/network"
)

func GetCars(ctx context.Context, carsCount int) ([]dto.SearchResponse, error) {
	iterations := calculateIterations(carsCount)
	result := make([]dto.SearchResponse, 0, iterations)

	for pageIndex := 1; pageIndex <= iterations; pageIndex++ {
		currentSize := calculateCarsOnPage(pageSize, pageIndex, iterations, carsCount)

		url, err := BuildURL(currentSize, pageIndex)
		if err != nil {
			return nil, fmt.Errorf("build url, page %d: %w", pageIndex, err)
		}

		resp, err := fetchPage(ctx, url)
		if err != nil {
			return nil, fmt.Errorf("fetch page %d: %w", pageIndex, err)
		}

		if resp.Result == nil {
			return nil, fmt.Errorf("empty result, page %d", pageIndex)
		}
		if resp.Result.TotalCount < carsCount {
			return nil, fmt.Errorf("page %d: %w", pageIndex, ErrCarsLessThenPassed)
		}

		result = append(result, *resp)
	}

	return result, nil
}

func fetchPage(ctx context.Context, url string) (*dto.SearchResponse, error) {
	body, err := network.MakeRequest(ctx, url, maxAttempts)
	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	resp, err := decodeSearchResponse(body)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return resp, nil
}
