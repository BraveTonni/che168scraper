package main

import (
	"flag"
	"path/filepath"
	"scraper/internal/scraper"
)

func main() {
	count := *flag.Int("count", 10, "car parsing count")
	outFileName := flag.String("output", "default", "output file name in folder out/ without .json")

	flag.Parse()

	output := filepath.Join("out/", *outFileName+".json")

	// Run scrapping from site
	scraper.Scrap(count, output)
}
