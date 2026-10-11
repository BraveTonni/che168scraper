package scraper

import "errors"

var ErrCarsLessThenPassed = errors.New("available cars on site less then passed by user")
