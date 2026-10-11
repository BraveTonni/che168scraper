package scraper

import (
	"log"
	"scraper/internal/logger"
)

func Scrap(count int, output string) {
	responses, err := GetCars(count)

	if err != nil {
		log.Fatal("Scrap: ", err)
	}

	for _, r := range responses {
		logger.Info(r.Result.CarList[len(r.Result.CarList)-1].BrandName)
	}

	logger.Info("Scrap finished. Result writed to: ", output)
}
