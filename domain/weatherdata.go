package domain

import "time"

type WeatherData struct {
	ID                 string    `json:"id"`
	Latitude           float32   `json:"latitude"`
	Longitude          float32   `json:"longitude"`
	Time               time.Time `json:"time"`
	WindSpeed          float32   `json:"windSpeed"`
	WindDirection      float32   `json:"windDirection"`
	AtmosphereTemp     float32   `json:"atmosphereTemp"`
	AtmospherePressure float32   `json:"atmospherePressure"`
	Humidity           float32   `json:"humidity"`
}

type WeatherDataCollection struct {
	Wdc []WeatherData
}

func (wdc *WeatherDataCollection) AddData(data WeatherData) []WeatherData {
	wdc.Wdc = append(wdc.Wdc, data)
	return wdc.Wdc

}
