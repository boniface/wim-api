package services

import (
	"fmt"
	"github.com/sivaramalingamk/regression"
	"wim-api/domain"
	"wim-api/ml"
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

func ProcessOutputData(id string, year int) (domain.OutputDataCollection, error) {
	odc := domain.OutputDataCollection{}

	odc, err := repository.SelectOutputData(id, year)
	if err != nil {
		fmt.Println("Error in Processing OUTPUT Data")
		return odc, err
	}
	fmt.Println("From Service length=", len(odc.Odc))
	return odc, nil

}

func ProcessInference() *regression.Regression {
	tdc, m := repository.SelectAllTraining()
	//r regression.Regression{}
	if m == "success" {
		r := ml.Regression(tdc)
		return r
	} else {
		println("Error while selecting trainig data")
		return nil
	}
	//return r
}
