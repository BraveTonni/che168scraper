package scraper_test

import (
	"scraper/internal/scraper"
	"testing"
)

func TestCalculateIterations(t *testing.T) {
	tests := []struct {
		count    int
		expected int
	}{
		{count: 10, expected: 1},
		{count: 100, expected: 2},
		{count: 101, expected: 3},
		{count: 0, expected: 0},
		{count: -50, expected: 0},
	}

	for i, tt := range tests {
		got := scraper.CalculateIterations(tt.count)

		if got != tt.expected {
			t.Errorf("test calculate iterations: %d. got: %d, expected: %d\n", i, got, tt.expected)
		}
	}
}

func TestCalculateCarsOnPage(t *testing.T) {
	tests := []struct {
		pageSize   int
		pageIndex  int
		iterations int
		carsCount  int
		expected   int
	}{
		// Last page, carsCount is evenly divisible by pageSize
		{pageSize: 50, pageIndex: 3, iterations: 3, carsCount: 150, expected: 50},
		// Last page, there is a remainder
		{pageSize: 50, pageIndex: 3, iterations: 3, carsCount: 105, expected: 5},
		// Last page, remainder equals 1
		{pageSize: 50, pageIndex: 3, iterations: 3, carsCount: 101, expected: 1},
		// Not the last page (first page)
		{pageSize: 50, pageIndex: 1, iterations: 3, carsCount: 105, expected: 50},
		// Not the last page (middle page)
		{pageSize: 50, pageIndex: 2, iterations: 3, carsCount: 105, expected: 50},
		// Single page, carsCount is less than pageSize
		{pageSize: 50, pageIndex: 1, iterations: 1, carsCount: 30, expected: 30},
		// Single page, carsCount equals pageSize
		{pageSize: 50, pageIndex: 1, iterations: 1, carsCount: 50, expected: 50},
		// pageSize = 1
		{pageSize: 1, pageIndex: 5, iterations: 5, carsCount: 5, expected: 1},
	}

	for i, tt := range tests {
		got := scraper.CalculateCarsOnPage(tt.pageSize, tt.pageIndex, tt.iterations, tt.carsCount)

		if got != tt.expected {
			t.Errorf("test %d (pageSize=%d, pageIndex=%d, iterations=%d, carsCount=%d): got: %d, expected: %d\n",
				i, tt.pageSize, tt.pageIndex, tt.iterations, tt.carsCount, got, tt.expected)
		}
	}
}
