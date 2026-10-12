package network_test

import (
	"testing"
)

func TestMakeRequest(t *testing.T) {

	tests := []struct {
		name    string
		fixture string
		url     string
		retries int
	}{
		{name: "0 cars in response", fixture: "cars_0.json", url: "", retries: 5},
		{name: "1 car in response", fixture: "cars_1.json", url: "", retries: 5},
		{name: "50 cars in response", fixture: "cars_50.json", url: "", retries: 5},
		{name: "101 car in response", fixture: "cars_101.json", url: "", retries: 5},
	}
	_ = tests
	// TODO:
	// data, err := os.ReadFile("testdata/cars_0.json")
	// if err != nil {
	// 	t.Fatalf("fail to load test data cars_0.json: %v", err)
	// }
	// _ = data
}
