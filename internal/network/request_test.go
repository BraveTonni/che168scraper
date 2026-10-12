package network_test

import (
	"context"
	"scraper/internal/network"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Arrange
	url := "https://globalapi.che168.com/api/v1/search?_appid=global.pc&deviceid=447c782c-b15f-478f-b928-85d4f8d7f74a&language=ru&fromsource=PCshouyeyoushangjia&pageindex=2&pagesize=24&sort=0&vehicle_list=1&report_status=1,3"

	// Act
	_, err := network.MakeRequest(ctx, url, 5)

	// Assert
	if err != nil {
		t.Error("make request: %w", err)
	}
}
