package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"scraper/internal/dto"
	"scraper/internal/logger"
	"scraper/internal/network"
	"strconv"
)

func GetCars(ctx context.Context, count int) ([]dto.SearchResponse, error) {
	iterations := count / pageSize
	// We add one to the number of iterations, because if the user passes a number that is not divisible by 50 without a remainder, then we will lose the remaining part.
	if count%pageSize != 0 {
		iterations++
	}

	var result []dto.SearchResponse

	for pageIndex := range iterations {
		pageIndex++

		u, err := url.Parse(baseURL)

		if err != nil {
			return result, fmt.Errorf("get cars: %w", err)
		}

		q := u.Query()
		q.Set("pagesize", strconv.Itoa(pageSize))
		q.Set("pageindex", strconv.Itoa(pageIndex))
		u.RawQuery = q.Encode()

		resp, err := network.MakeRequest(ctx, u.String())
		if err != nil {
			// We do not return an error, as this may be a one-off
			logger.Warn("Failed to get cars from page: ", err)
			continue
		}

		var searchResponse dto.SearchResponse
		if err = json.Unmarshal(resp, &searchResponse); err != nil {
			logger.Warn("Failed to unmarshal json: ", err)
			continue
		}

		// If the user entered a number of cars greater than the number on the website, we return an error
		if searchResponse.Result.TotalCount < count {
			return nil, fmt.Errorf("get cars: %w", ErrCarsLessThenPassed)
		}

		// Last page
		if pageIndex == iterations {
			// Without verification, the page will be truncated to zero.
			if count%pageSize != 0 {
				remainedCars := count % pageSize
				searchResponse.Result.CarList = searchResponse.Result.CarList[:remainedCars]
				searchResponse.Result.CarsCount = remainedCars
			}

			logger.Debugf("Scrapped %d/%d cars\n", count, count)
		} else {
			logger.Debugf("Scrapped %d/%d cars\n", pageSize*pageIndex, count)
		}

		result = append(result, searchResponse)
	}

	return result, nil
}
