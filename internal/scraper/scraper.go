package scraper

import (
	"context"
	"fmt"
	"scraper/internal/fileutil"
)

func Scrap(ctx context.Context, count int, output string) error {
	responses, err := GetCars(ctx, count)

	if err != nil || len(responses) == 0 {
		return fmt.Errorf("scrap: %w", err)
	}

	err = fileutil.WriteResponse(output, responses)

	if err != nil {
		return fmt.Errorf("scrap: %w", err)
	}

	return nil
}
