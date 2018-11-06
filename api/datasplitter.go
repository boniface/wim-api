package api

import "wim-api/domain"

func splitData(data domain.RawInputData) (domain.VehicleData, domain.Coordinate) {
	var vehicleData domain.VehicleData
	var coordinate domain.Coordinate

	vehicleData.ID = data.ID
	vehicleData.Time = data.CollectedTime
	vehicleData.VehicleSpeed = data.VehicleSpeed
	vehicleData.Acceleration = data.Acceleration
	vehicleData.CoolentTemp = data.CoolentTemp
	vehicleData.ElevationAngle = data.ElevationAngle
	vehicleData.EngineLoad = data.EngineLoad
	vehicleData.FuelFlow = data.FuelFlow
	vehicleData.HeadingDirection = data.HeadingDirection
	vehicleData.IntakeAirTemp = data.IntakeAirTemp
	vehicleData.O2 = data.O2
	vehicleData.Rpm = data.Rpm
	vehicleData.Weight = data.Weight
	vehicleData.OilPressure = data.OilPressure

	coordinate.Latitude = data.Latitude
	coordinate.Longitude = data.Longitude

	return vehicleData, coordinate

}

func splitBulkData(data []domain.RawInputData) (domain.VehicleDataCollection, domain.CoordinateCollection) {

	vdc := domain.VehicleDataCollection{}
	cc := domain.CoordinateCollection{}

	for i := range data {
		var vehicleData domain.VehicleData
		var coordinate domain.Coordinate

		vehicleData.ID = data[i].ID
		vehicleData.VehicleSpeed = data[i].VehicleSpeed
		vehicleData.Acceleration = data[i].Acceleration
		vehicleData.CoolentTemp = data[i].CoolentTemp
		vehicleData.ElevationAngle = data[i].ElevationAngle
		vehicleData.EngineLoad = data[i].EngineLoad
		vehicleData.FuelFlow = data[i].FuelFlow
		vehicleData.HeadingDirection = data[i].HeadingDirection
		vehicleData.IntakeAirTemp = data[i].IntakeAirTemp
		vehicleData.O2 = data[i].O2
		vehicleData.Rpm = data[i].Rpm
		vehicleData.Weight = data[i].Weight
		vehicleData.OilPressure = data[i].OilPressure
		vdc.AddData(vehicleData)

		coordinate.ID = data[i].ID
		coordinate.Latitude = data[i].Latitude
		coordinate.Longitude = data[i].Longitude
		cc.AddData(coordinate)
	}
	return vdc, cc

}
