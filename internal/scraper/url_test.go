package scraper_test

import (
	"errors"
	"scraper/internal/scraper"
	"testing"
)

func TestBuildURL(t *testing.T) {
	// Arrange
	tests := []struct {
		pageSize  int
		pageIndex int
		expected  string
		err       error
	}{
		{
			pageSize:  0,
			pageIndex: -1,
			expected:  "",
			err:       scraper.ErrInvalidParameters,
		},
		{
			pageSize:  -1,
			pageIndex: 0,
			expected:  "",
			err:       scraper.ErrInvalidParameters,
		},
		{
			pageSize:  -1,
			pageIndex: -1,
			expected:  "",
			err:       scraper.ErrInvalidParameters,
		},
		{
			pageSize:  0,
			pageIndex: 0,
			expected:  "",
			err:       scraper.ErrInvalidParameters,
		},
		{
			pageSize:  50,
			pageIndex: 2,
			expected:  "https://globalapi.che168.com/api/v1/search?_appid=global.pc&deviceid=447c782c-b15f-478f-b928-85d4f8d7f74a&fromsource=PCshouyeyoushangjia&language=ru&pageindex=2&pagesize=50&report_status=1%2C3&sort=0&vehicle_list=1",
			err:       nil,
		},
	}

	for i, tt := range tests {
		// Act
		got, err := scraper.BuildURL(tt.pageSize, tt.pageIndex)

		// Assert
		if got != tt.expected {
			t.Errorf("test build url: %d. got: %s, expected: %s\n", i, got, tt.expected)
		}

		if !errors.Is(err, tt.err) {
			t.Errorf("test build url: %d. got err: %v, expected err: %v\n", i, err, tt.err)
		}
	}

}
