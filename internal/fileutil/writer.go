package fileutil

import (
	"encoding/json"
	"os"
	"scraper/internal/logger"
	"scraper/internal/model"
)

func WriteResponse(out string, responses []model.SearchResponse) error {
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
