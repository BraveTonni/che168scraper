//go:build integration

package network_test

import (
	"context"
	"encoding/json"
	"scraper/internal/dto"
	"scraper/internal/network"
	"scraper/internal/scraper"
	"testing"
)

func TestMakeRequest_Integration(t *testing.T) {
	tests := []struct {
		name        string
		pageSize    int
		pageIndex   int
		wantMaxCars int // максимум = pageSize, но может быть меньше на последней странице
	}{
		{name: "pageSize 1", pageSize: 1, pageIndex: 1, wantMaxCars: 1},
		{name: "pageSize 50", pageSize: 50, pageIndex: 1, wantMaxCars: 50},
		{name: "pageSize 50, last page", pageSize: 50, pageIndex: 97, wantMaxCars: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := scraper.BuildURL(tt.pageSize, tt.pageIndex)
			if err != nil {
				t.Fatalf("BuildURL: %v", err)
			}

			resp, err := network.MakeRequest(context.Background(), url, 5)
			if err != nil {
				t.Fatalf("MakeRequest: %v", err)
			}

			var got dto.SearchResponse
			if err := json.Unmarshal(resp, &got); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}

			// --- инварианты, верные для любого ответа ---
			if got.ReturnCode != 0 {
				t.Errorf("returncode = %d, want 0", got.ReturnCode)
			}
			if got.Result.PageSize != tt.pageSize {
				t.Errorf("pagesize = %d, want %d", got.Result.PageSize, tt.pageSize)
			}
			if got.Result.PageIndex != tt.pageIndex {
				t.Errorf("pageindex = %d, want %d", got.Result.PageIndex, tt.pageIndex)
			}
			if len(got.Result.CarList) == 0 {
				t.Error("carlist is empty")
			}
			if len(got.Result.CarList) > tt.wantMaxCars {
				t.Errorf("carlist length = %d, want <= %d", len(got.Result.CarList), tt.wantMaxCars)
			}
			if got.Result.TotalCount <= 0 {
				t.Errorf("totalcount = %d, want > 0", got.Result.TotalCount)
			}
			if got.Result.CarsCount <= 0 {
				t.Errorf("pagecount = %d, want > 0", got.Result.CarsCount)
			}
		})
	}
}
