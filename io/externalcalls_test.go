package io

import (
	"fmt"
	"testing"
	"wim-api/domain"
)

func TestWeatherAPI(t *testing.T) {
	data := domain.Coordinate{
		Latitude:  20.43,
		Longitude: 60.567,
	}
	fmt.Println("Sending coordinate  Points : ", data.Longitude, data.Latitude)
	result, _ := WeatherAPI(data)
	fmt.Println("The Result is  ", result)

}
