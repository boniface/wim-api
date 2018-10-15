package services

import (
	"wim-api/domain"
	"wim-api/repository"
)

func ProcessVehicleData(data domain.VehicleData) (string, error) {

	if _, err := repository.AddVehicleData(data); err != nil {
		println("Error in Processing Vehicle Data")
		return "", err

	}

	return "Success", nil

}

func ProcessVehicleDataCollection(data domain.VehicleDataCollection) (string, error) {

	if _, err := repository.AddVehicleDataCollection(data); err != nil {
		println("Error in Processing Vehicle Data Collection")
		return "", err

	}

	return "Success", nil

}

func ProcessWetherData(data domain.WeatherData) (string, error) {

	if _, err := repository.AddWeatherData(data); err != nil {
		println("Error in Processing Weather Data")
		return "", err

	}

	return "Success", nil
}

func ProcessBulkWetherData(data domain.WeatherDataCollection) (string, error) {

	if _, err := repository.AddWeatherDataCollection(data); err != nil {
		println("Error in Processing Weather Data")
		return "", err

	}

	return "Success", nil
}

func ProcessTrainingData(data domain.TrainingDataCollection) (string, error) {
	if len(data.Tdc) > 1 {
		if _, err := repository.AddTrainingData(data); err != nil {
			println("Error in Processing Training Data")
			return "", err

		}
	}
	return "Success", nil

}
