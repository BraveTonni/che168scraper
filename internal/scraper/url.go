package scraper

import (
	"fmt"
	"net/url"
	"strconv"
)

func BuildURL(PageSize int, PageIndex int) (string, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return "", fmt.Errorf("parse url: %w", err)
	}

	if PageSize <= 0 || PageIndex <= 0 {
		return "", fmt.Errorf("parse url: %w", ErrInvalidParameters)
	}

	q := u.Query()
	q.Set("pagesize", strconv.Itoa(PageSize))
	q.Set("pageindex", strconv.Itoa(PageIndex))
	u.RawQuery = q.Encode()

	return u.String(), nil
}
