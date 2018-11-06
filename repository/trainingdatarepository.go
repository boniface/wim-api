package repository

import (
	"fmt"
	"wim-api/domain"
)

func AddTrainingData(tdc domain.TrainingDataCollection) (string, error) {
	fmt.Println(" **** Creating new Training data****\n Adding ", len(tdc.Tdc), "Number of Training Data")
	defer getCluster().Close()
	for _, td := range tdc.Tdc {
		if err := Session.Query("INSERT INTO trainingdata( id,acceleration, atmospherepressure, atmospheretemp, coolanttemp, elevationangle, engineload, fuelflow, headingdirection, humidity, intakeairtemp, o2, oilpressure, rpm, vehiclespeed, weight, winddirection, windspeed ) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			td.ID, td.Acceleration, td.AtmospherePressure, td.AtmosphereTemp, td.CoolantTemp, td.ElevationAngle, td.EngineLoad, td.FuelFlow, td.HeadingDirection, td.Humidity, td.IntakeAirTemp, td.O2, td.OilPressure, td.Rpm, td.VehicleSpeed, td.Weight, td.WindDirection, td.WindSpeed).Exec(); err != nil {

			fmt.Println("Error while inserting Training Data", err)
			fmt.Println(err)
			return "", err
		}
	}
	return "Success", nil
}

func SelectAllTraining() (domain.TrainingDataCollection, string) {
	fmt.Println("Selecting All Training data  ")
	defer getCluster().Close()
	temp := domain.TrainingData{}
	res := domain.TrainingDataCollection{}
	iter := Session.Query(`SELECT id,acceleration,atmospherepressure, atmospheretemp,coolanttemp,elevationangle,engineload,fuelflow,headingdirection, humidity, intakeairtemp , o2 ,oilpressure ,rpm ,vehiclespeed ,weight ,winddirection ,windspeed
FROM trainingdata`).Iter()
	for iter.Scan(&temp.ID, &temp.Acceleration, &temp.AtmospherePressure, &temp.AtmosphereTemp, &temp.CoolantTemp, &temp.ElevationAngle, &temp.EngineLoad, &temp.FuelFlow, &temp.HeadingDirection, &temp.Humidity, &temp.IntakeAirTemp, &temp.O2, &temp.OilPressure, &temp.Rpm, &temp.VehicleSpeed, &temp.Weight, &temp.WindDirection, &temp.WindSpeed) {
		res.AddData(temp)
	}

	if err := iter.Close(); err != nil {
		return res, "Error"
	}
	return res, "success"
}
