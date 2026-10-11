package scraper

import (
	"scraper/internal/fileutil"
	"scraper/internal/logger"
)

func Scrap(count int, output string) {
	responses, err := GetCars(count)

	if err != nil || len(*responses) == 0 {
		logger.Fatal("Scrap: ", err)
	}

	err = fileutil.WriteResponse(output, *responses)

	if err != nil {
		logger.Error("Fail write response to file: ", err)
	}

	logger.Info("Scrap finished. Result writed to: ", output)
}
