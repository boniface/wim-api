package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wim-api/domain"
)

func WeatherAPI(coordinate domain.Coordinate) (domain.WeatherData, error) {
	baseUrl := "https://api.openweathermap.org/data/2.5/weather"
	appid := API_Key
	url1 := fmt.Sprintf(baseUrl+"?lat=%f&lon=%f&appid=%s", coordinate.Latitude, coordinate.Longitude, appid)
	fmt.Println("URL:>", url1)

	req, err := http.NewRequest("GET", url1, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	owmdata := domain.OwmApiData{}
	err1 := json.Unmarshal(body, &owmdata)
	if err != nil {
		panic(err1)
	}
	wdata := owmDataToWeather(owmdata)
	fmt.Println("Temp=", wdata.AtmosphereTemp)
	wdata.ID = coordinate.ID
	wdata.Latitude = coordinate.Latitude
	wdata.Longitude = coordinate.Longitude
	return wdata, err
}

func owmDataToWeather(data domain.OwmApiData) domain.WeatherData {
	var wd domain.WeatherData

	wd.AtmosphereTemp = float32(data.Main.Temp)
	wd.Humidity = float32(data.Main.Humidity)
	wd.AtmospherePressure = float32(data.Main.Pressure)
	wd.WindDirection = float32(data.Wind.Deg)
	wd.WindSpeed = float32(data.Wind.Speed)

	return wd
}
