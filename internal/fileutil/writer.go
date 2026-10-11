package fileutil

import (
	"encoding/json"
	"os"
	"scraper/internal/logger"
	"scraper/internal/model"
)

func WriteResponse(out string, responses []model.SearchResponse) {
	data, err := json.MarshalIndent(responses, "", jsonIndent)

	if err != nil {
		logger.Error("json serializition: ", err)
	}

	_, err = os.Stat(out)

	if err == nil {
		logger.Warnf("File %s rewrited\n", out)
	}

	err = os.WriteFile(out, data, 0755)

	if err != nil {
		logger.Error("write bytes to file: ", err)
	}
}
