package fileutil

import (
	"encoding/json"
	"log"
	"os"
	"scraper/internal/logger"
	"scraper/internal/model"
)

func WriteResponse(out string, responses []model.SearchResponse) {
	data, err := json.MarshalIndent(responses, "", jsonIndent)

	if err != nil {
		log.Fatal("json serializition: ", err)
	}

	_, err = os.Stat(out)

	if err == nil {
		logger.Warnf("File %s rewrited\n", out)
	}

	err = os.WriteFile(out, data, 0644)

	if err != nil {
		log.Fatal("write bytes to file: ", err)
	}
}
