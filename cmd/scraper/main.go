package main

import (
	"flag"
	"path/filepath"
	"scraper/internal/logger"
	"scraper/internal/scraper"
)

func main() {
	logger.Init()

	count := flag.Int("count", 15, "car parsing count")
	outFileName := flag.String("output", "default", "output file name in folder out/ without .json")

	flag.Parse()

	if *count <= 0 {
		logger.Fatal("Count must be > 0.")
	}

	output := filepath.Join("out/", *outFileName+".json")

	logger.Info("Start scrapping cars")

	// Run scrapping from site
	scraper.Scrap(*count, output)
}
