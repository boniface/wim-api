package repository

import (
	"fmt"
	"log"
	"wim-api/domain"
)

func AddWeatherData(data domain.WeatherData) (string, error) {
	fmt.Println(" **** Creating new Weather data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO weatherdata(id,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed) values(?,?,?,?,?,?)", data.ID, data.AtmospherePressure, data.AtmosphereTemp, data.Humidity, data.WindDirection, data.WindSpeed).Exec(); err != nil {

		fmt.Println("Error while inserting Weather Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}

func AddWeatherDataCollection(wdc domain.WeatherDataCollection) (string, error) {
	fmt.Println(" **** Creating new Weather data ****\n Inserting ", len(wdc.Wdc), "number of rows")
	defer getCluster().Close()
	for _, data := range wdc.Wdc {

		if err := Session.Query("INSERT INTO weatherdata(id,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed) values(?,?,?,?,?,?)", data.ID, data.AtmospherePressure, data.AtmosphereTemp, data.Humidity, data.WindDirection, data.WindSpeed).Exec(); err != nil {
			fmt.Println("Error while inserting Weather Data", err)
			fmt.Println(err)
			return "", err
		}

	}
	return "Success", nil
}

func SelectWeatherData() (domain.WeatherDataCollection, string) {
	fmt.Println("Selecting Weather Data colection From Database")
	defer getCluster().Close()
	wdata := []domain.WeatherData{}
	wDataCollection := domain.WeatherDataCollection{wdata}
	temp := domain.WeatherData{}
	iter := Session.Query(`SELECT id,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed FROM weatherdata`).Iter()
	for iter.Scan(&temp.ID, &temp.AtmospherePressure, &temp.AtmosphereTemp, &temp.Humidity, &temp.WindDirection, &temp.WindSpeed) {
		wDataCollection.AddData(temp)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return wDataCollection, "Error"
	}
	fmt.Println("Selected ", len(wDataCollection.Wdc), " Number of Weather Data")
	return wDataCollection, "success"
}
