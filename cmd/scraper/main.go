package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"scraper/internal/logger"
	"scraper/internal/scraper"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger.Init()

	count := flag.Int("count", 15, "car parsing count")
	outFileName := flag.String("output", "default", "output file name in folder out/ without .json")

	flag.Parse()

	if *count <= 0 {
		logger.Fatal("Count must be > 0.")
	}

	os.MkdirAll("out", 0o755)
	output := filepath.Join("out/", *outFileName+".json")

	logger.Info("Start scrapping cars")

	// Run scrapping from site
	err := scraper.Scrap(ctx, *count, output)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Scrap finished. Result writed to: ", output)
}
