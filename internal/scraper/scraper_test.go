package scraper_test

import (
	"context"
	"errors"
	"scraper/internal/scraper"
	"testing"
)

func TestScrap(t *testing.T) {
	ctx := context.Background()
	err := scraper.Scrap(ctx, 0, "")

	if !errors.Is(err, scraper.ErrNoCars) {
		t.Errorf("test scrap: %v", err)
	}

	err = scraper.Scrap(ctx, 10, "")

	if err != nil {
		t.Errorf("test scrap count 10: %v", err)
	}

	err = scraper.Scrap(ctx, 51, "")

	if err != nil {
		t.Errorf("test scrap count 51: %v", err)
	}
}
