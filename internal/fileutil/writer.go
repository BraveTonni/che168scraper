package fileutil

import (
	"encoding/json"
	"os"
	"scraper/internal/dto"
	"scraper/internal/logger"
)

func WriteResponse(out string, responses []dto.SearchResponse) error {
	data, err := json.MarshalIndent(responses, "", jsonIndent)

	if err != nil {
		return err
	}

	_, err = os.Stat(out)

	if err == nil {
		logger.Warnf("File %s rewrited\n", out)
	}

	err = os.WriteFile(out, data, 0755)

	if err != nil {
		return err
	}

	return nil
}
