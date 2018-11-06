package domain

import (
	"errors"
	"net/http"
	"time"
)

type VehicleData struct {
	ID               string    `json:"id"`
	Time             time.Time `json:"time"`
	Weight           float32   `json:"weight"`
	VehicleSpeed     float32   `json:"vehicleSpeed"`
	Acceleration     float32   `json:"acceleration"`
	HeadingDirection float32   `json:"headingDirection"`
	CoolentTemp      float32   `json:"coolentTemp"`
	OilPressure      float32   `json:"oilPressure"`
	IntakeAirTemp    float32   `json:"intake_airTemp"`
	Rpm              float32   `json:"rpm"`
	EngineLoad       float32   `json:"engineLoad"`
	ElevationAngle   float32   `json:"elevationAngle"`
	O2               float32   `json:"o2"`
	FuelFlow         float32   `json:"fuelFlow"`
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
