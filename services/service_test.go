package services

import (
	"fmt"
	"testing"
	"wim-api/domain"
)

func TestAddVehicle(t *testing.T) {
	data := domain.VehicleData{
		ID:               "12322",
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      50,
		OilPressure:      69,
		IntakeAirTemp:    70,
		Rpm:              80,
		EngineLoad:       90,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}

	msg, err := ProcessVehicleData(data)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("The Result  ", msg)

}

func TestAddWeather(t *testing.T) {
	data := domain.WeatherData{
		ID:                 "12322",
		WindSpeed:          40,
		WindDirection:      45,
		AtmosphereTemp:     200,
		AtmospherePressure: 3,
		Humidity:           94,
	}

	msg, err := ProcessWetherData(data)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("The Result  ", msg)

}
