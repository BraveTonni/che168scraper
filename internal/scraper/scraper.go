package scraper

import (
	"log"
	"scraper/internal/fileutil"
	"scraper/internal/logger"
)

func Scrap(count int, output string) {
	responses, err := GetCars(count)

	if err != nil || len(responses) == 0 {
		log.Fatal("Scrap: ", err)
	}

	fileutil.WriteResponse(output, responses)

	logger.Info("Scrap finished. Result writed to: ", output)
}
