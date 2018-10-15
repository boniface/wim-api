package api

import (
	"fmt"
	"github.com/robfig/cron"
	"wim-api/domain"
	"wim-api/repository"
	"wim-api/services"
)

func MergeAndInsertTraining() {
	//	data:=[]domain.TrainingData{}
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		fmt.Println("Testing Cron")
		tdc, _ := merger()
		services.ProcessTrainingData(tdc)

	})
	c.Start()

}

func merger() (domain.TrainingDataCollection, string) {

	wdc, _ := repository.SelectWeatherData()
	vdc, _ := repository.SelectAllVehicle()
	tdc := domain.TrainingDataCollection{}
	for _, wd := range wdc.Wdc {
		for _, vd := range vdc.Vdc {
			if vd.ID == wd.ID {
				fmt.Println("Adding values VehicleID:", vd.ID, "WeatherID:", wd.ID)
				tdc.AddData(trainDataMerger(vd, wd))
			}
		}

	}

	return tdc, "Success"
}

func trainDataMerger(vdata domain.VehicleData, wdata domain.WeatherData) domain.TrainingData {
	var tdata domain.TrainingData

	tdata.ID = vdata.ID
	tdata.AtmospherePressure = wdata.AtmospherePressure
	tdata.WindSpeed = wdata.WindSpeed
	tdata.WindDirection = wdata.WindDirection
	tdata.Humidity = wdata.Humidity
	tdata.AtmosphereTemp = wdata.AtmosphereTemp
	tdata.FuelFlow = vdata.FuelFlow
	tdata.O2 = vdata.O2
	tdata.ElevationAngle = vdata.ElevationAngle
	tdata.EngineLoad = vdata.EngineLoad
	tdata.Rpm = vdata.Rpm
	tdata.IntakeAirTemp = vdata.IntakeAirTemp
	tdata.OilPressure = vdata.OilPressure
	tdata.CoolantTemp = vdata.CoolentTemp
	tdata.HeadingDirection = vdata.HeadingDirection
	tdata.Acceleration = vdata.Acceleration
	tdata.Weight = vdata.Weight
	tdata.VehicleSpeed = vdata.VehicleSpeed
	return tdata
}
