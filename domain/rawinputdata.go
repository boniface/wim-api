package domain

import (
	"errors"
	"net/http"
)

type RawInputData struct {
	ID               string  `json:"id"`
	Time             string  `json:"time"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Weight           int     `json:"weight"`
	VehicleSpeed     int     `json:"vehicleSpeed"`
	Acceleration     int     `json:"acceleration"`
	HeadingDirection int     `json:"headingDirection"`
	CoolentTemp      int     `json:"coolentTemp"`
	OilPressure      int     `json:"oilPressure"`
	IntakeAirTemp    int     `json:"intakeAirTemp"`
	Rpm              int     `json:"rpm"`
	EngineLoad       int     `json:"engineLoad"`
	ElevationAngle   int     `json:"elevationAngle"`
	O2               int     `json:"o2"`
	FuelFlow         int     `json:"fuelFlow"`
}

func (vdata *RawInputData) Bind(r *http.Request) error {
	if vdata.ID == "" {
		return errors.New("missing required RawInputData fields.")
	}

	return nil
}

type RawInputDataCollection struct {
	Idc []RawInputData
}

func (idc *RawInputDataCollection) AddData(data RawInputData) []RawInputData {
	idc.Idc = append(idc.Idc, data)
	return idc.Idc

}
