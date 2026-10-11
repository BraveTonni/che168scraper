package model

import "time"

type Car struct {
	Name     string
	SpecName string
	ImageURL string
	Price    float32

	FuelName         string
	RegistrationDate time.Time
	Mileage          int
	Transmission     int

	PhoneNumber int

	WeChatURL   string
	WhatsappURL string
}
