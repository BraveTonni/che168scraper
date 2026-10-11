package scraper_test

import (
	"scraper/internal/scraper"
	"testing"
)

func TestGetCars(t *testing.T) {
	tests := []struct {
		name      string
		count     int
		wantPages int // count of cars on page
	}{
		{
			name:      "Zero count passed",
			count:     0,
			wantPages: 0,
		},
		{
			name:      "The number is divisible by 50 without remainder",
			count:     50,
			wantPages: 1,
		},
		{
			name:      "The number is divisible by 50 with remainder",
			count:     51,
			wantPages: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			responses, err := scraper.GetCars(tt.count)

			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(responses) != tt.wantPages {
				t.Errorf("got %d pages, want %d", len(responses), tt.wantPages)
			}
		})
	}
}
