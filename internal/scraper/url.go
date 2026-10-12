package scraper

import (
	"fmt"
	"net/url"
	"strconv"
)

func buildURL(pageSize int, pageIndex int) (string, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return "", fmt.Errorf("parse url: %w", err)
	}

	q := u.Query()
	q.Set("pagesize", strconv.Itoa(pageSize))
	q.Set("pageindex", strconv.Itoa(pageIndex))
	u.RawQuery = q.Encode()

	return u.String(), nil
}
