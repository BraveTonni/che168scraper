package network_test

import (
	"scraper/internal/network"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	// Arrange
	url := "https://globalapi.che168.com/api/v1/search?_appid=global.pc&deviceid=447c782c-b15f-478f-b928-85d4f8d7f74a&language=ru&fromsource=PCshouyeyoushangjia&pageindex=2&pagesize=24&sort=0&vehicle_list=1&report_status=1,3"

	// Act
	_, err := network.MakeRequest(url)

	// Assert
	if err != nil {
		t.Error("make request: %w", err)
	}
}

func TestMakeRequestError(t *testing.T) {
	// Arrange
	url := "https://unexisted.lan/"

	// Act
	_, err := network.MakeRequest(url)

	// Assert
	if err == nil {
		t.Error("make request unexisted site: %w", err)
	}
}
