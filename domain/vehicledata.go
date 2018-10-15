package domain

import (
	"errors"
	"net/http"
)

type VehicleData struct {
	ID               string `json:"id"`
	Weight           int    `json:"weight"`
	VehicleSpeed     int    `json:"vehicleSpeed"`
	Acceleration     int    `json:"acceleration"`
	HeadingDirection int    `json:"headingDirection"`
	CoolentTemp      int    `json:"coolentTemp"`
	OilPressure      int    `json:"oilPressure"`
	IntakeAirTemp    int    `json:"intake_airTemp"`
	Rpm              int    `json:"rpm"`
	EngineLoad       int    `json:"engineLoad"`
	ElevationAngle   int    `json:"elevationAngle"`
	O2               int    `json:"o2"`
	FuelFlow         int    `json:"fuelFlow"`
}

type VehicleDataCollection struct {
	Vdc []VehicleData
}

func (vdc *VehicleDataCollection) AddData(data VehicleData) []VehicleData {
	vdc.Vdc = append(vdc.Vdc, data)
	return vdc.Vdc
}

func (vdata *VehicleData) Bind(r *http.Request) error {
	if vdata.ID == "" {
		return errors.New("missing required VehicleData fields.")
	}

	return nil
}
