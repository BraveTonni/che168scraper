# che168scraper

A tool for scraping a specified number of cars from the che168 website (https://global.che168.com/)

## How to use?
* Windows: scraper.exe -count 50 -output cars
* Linux: ./scraper -count 50 -output cars

## How it works?
* Sends a request to the site's hidden API
* Parses the JSON response
* Writes the collected cars to the output file

## RunBook: how to start?
* make up - command to run (for development)
* make test - command to run tests (for development)

## Requirements
* go 1.27+
* Make (optional, for the RunBook commands)
* Internet access to global.che168.com