package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
	"testing"
	"time"
	"wim-api/domain"
)

func TestHomePage(t *testing.T) {
	resp, _ := resty.R().Get("http://localhost:3000")
	assert.Equal(t, 200, resp.StatusCode())
}

func TestSimpleRawData(t *testing.T) {
	data := domain.RawInputData{
		ID:               "12322",
		CollectedTime:    time.Now(),
		Latitude:         122,
		Longitude:        15,
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

	resp, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post("http://localhost:3000/ecudata/simple")

	fmt.Println("The Result For Single Data Point is ", resp)
	fmt.Println("The Result For Single Data Point  ", resp.Status())
	assert.Equal(t, "200 OK", resp.Status())
}

func TestIncomingRawDataCollection(t *testing.T) {
	data1 := domain.RawInputData{
		ID:               "sk223",
		CollectedTime:    time.Now(),
		Latitude:         48,
		Longitude:        95,
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     210,
		HeadingDirection: 401,
		CoolentTemp:      50,
		OilPressure:      691,
		IntakeAirTemp:    701,
		Rpm:              801,
		EngineLoad:       901,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}

	data2 := domain.RawInputData{
		ID:               "sk223",
		CollectedTime:    time.Now(),
		Latitude:         19,
		Longitude:        43,
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      50,
		OilPressure:      69,
		IntakeAirTemp:    70,
		Rpm:              801,
		EngineLoad:       90,
		ElevationAngle:   1010,
		O2:               201,
		FuelFlow:         30,
	}

	data3 := domain.RawInputData{
		ID:               "sk121",
		CollectedTime:    time.Now(),
		Latitude:         45,
		Longitude:        89,
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      510,
		OilPressure:      619,
		IntakeAirTemp:    710,
		Rpm:              810,
		EngineLoad:       90,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}
	data := []domain.RawInputData{data1, data2, data3}
	resp, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post("http://localhost:8080/ecudata/bulk")
	fmt.Println("The Result For Collections is ", resp)
	fmt.Println("The Result Collections is ", resp.Status())
	assert.Equal(t, "200 OK", resp.Status())
}
