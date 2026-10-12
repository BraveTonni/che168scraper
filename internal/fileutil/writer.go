package fileutil

import (
	"encoding/json"
	"fmt"
	"os"
	"scraper/internal/dto"
	"scraper/internal/logger"
)

func WriteResponse(out string, responses []dto.SearchResponse) error {
	data, err := json.MarshalIndent(responses, "", jsonIndent)

	if err != nil {
		return fmt.Errorf("write response to file: %w", err)
	}

	_, err = os.Stat(out)

	if err == nil {
		logger.Debugf("File %s rewrited\n", out)
	}

	err = os.WriteFile(out, data, 0o755)

	if err != nil {
		return fmt.Errorf("write response to file: %w", err)
	}

	return nil
}
