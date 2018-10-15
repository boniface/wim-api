package repository

import (
	"fmt"
	"github.com/gocql/gocql"
	"wim-api/domain"
)

func AddVehicleData(data domain.VehicleData) (string, error) {
	fmt.Println(" **** Creating new vehicle data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO vehicledata(id,weight,vehiclespeed,acceleration,headingdirection,coolenttemp,oilpressure,intakeairtemp,rpm,engineload,elevationangle,o2,fuelflow) values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
		data.ID, data.Weight, data.VehicleSpeed, data.Acceleration, data.HeadingDirection, data.CoolentTemp, data.OilPressure, data.IntakeAirTemp, data.Rpm, data.EngineLoad, data.ElevationAngle, data.O2, data.FuelFlow).Exec(); err != nil {

		fmt.Println("Error while inserting Vehicle Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}

func AddVehicleDataCollection(vdc domain.VehicleDataCollection) (string, error) {
	fmt.Println(" **** Creating new Vechicle data Collection ****\nAdding ", len(vdc.Vdc), " Number of Vehicle Data")
	defer getCluster().Close()
	for _, data := range vdc.Vdc {
		if err := Session.Query("INSERT INTO vehicledata(id,weight,vehiclespeed,acceleration,headingdirection,coolenttemp,oilpressure,intakeairtemp,rpm,engineload,elevationangle,o2,fuelflow) values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
			data.ID, data.Weight, data.VehicleSpeed, data.Acceleration, data.HeadingDirection, data.CoolentTemp, data.OilPressure, data.IntakeAirTemp, data.Rpm, data.EngineLoad, data.ElevationAngle, data.O2, data.FuelFlow).Exec(); err != nil {

			fmt.Println("Error while inserting Vehicle Data Collection", err)
			fmt.Println(err)
			return "", err
		}
	}
	return "Success", nil
}

func SelectVehicleData(id string) (domain.VehicleData, string) {
	fmt.Println("Selecting Vehicle data with ID : ", id)
	defer getCluster().Close()
	temp := domain.VehicleData{}
	err1 := Session.Query(`SELECT  id, acceleration, coolenttemp, elevationangle, engineload, fuelflow, headingdirection, intakeairtemp, o2, oilpressure, rpm, vehiclespeed, weight 
	FROM vehicledata WHERE id=? LIMIT 1`, id).Consistency(gocql.One).Scan(&temp.ID, &temp.Acceleration, &temp.CoolentTemp, &temp.ElevationAngle, &temp.EngineLoad, &temp.FuelFlow, &temp.HeadingDirection, &temp.IntakeAirTemp, &temp.O2, &temp.OilPressure, &temp.Rpm, &temp.VehicleSpeed, &temp.Weight)

	if err1 != nil {
		fmt.Println("ERror")
		return temp, "Error"
	}

	return temp, "success"
}

func SelectAllVehicle() (domain.VehicleDataCollection, string) {
	fmt.Println("Selecting All Vehicle data  ")
	defer getCluster().Close()
	temp := domain.VehicleData{}
	res := domain.VehicleDataCollection{}
	iter := Session.Query(`SELECT  id, acceleration, coolenttemp, elevationangle, engineload, fuelflow, headingdirection, intakeairtemp, o2, oilpressure, rpm, vehiclespeed, weight 
	FROM vehicledata`).Iter()
	for iter.Scan(&temp.ID, &temp.Acceleration, &temp.CoolentTemp, &temp.ElevationAngle, &temp.EngineLoad, &temp.FuelFlow, &temp.HeadingDirection, &temp.IntakeAirTemp, &temp.O2, &temp.OilPressure, &temp.Rpm, &temp.VehicleSpeed, &temp.Weight) {
		res.AddData(temp)
	}

	if err := iter.Close(); err != nil {
		return res, "Error"
	}
	return res, "success"
}
