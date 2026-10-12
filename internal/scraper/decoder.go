package scraper

import (
	"encoding/json"
	"fmt"
	"scraper/internal/dto"
	"scraper/internal/logger"
)

func decodeSearchResponse(bytes []byte) (*dto.SearchResponse, error) {
	var searchResponse dto.SearchResponse

	if err := json.Unmarshal(bytes, &searchResponse); err != nil {
		logger.Debug("Failed to unmarshal json: ", err)
		return nil, fmt.Errorf("decode search response, json unmarshal: %w", err)
	}

	return &searchResponse, nil
}
