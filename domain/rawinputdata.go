package domain

import (
	"errors"
	"net/http"
	"time"
)

type RawInputData struct {
	ID               string    `json:"id"`
	CollectedTime    time.Time `json:"collectedTime"`
	InsertedTime     time.Time `json:"insertedTime"`
	Latitude         float32   `json:"latitude"`
	Longitude        float32   `json:"longitude"`
	Weight           float32   `json:"weight"`
	VehicleSpeed     float32   `json:"vehicleSpeed"`
	Acceleration     float32   `json:"acceleration"`
	HeadingDirection float32   `json:"headingDirection"`
	CoolentTemp      float32   `json:"coolentTemp"`
	OilPressure      float32   `json:"oilPressure"`
	IntakeAirTemp    float32   `json:"intakeAirTemp"`
	Rpm              float32   `json:"rpm"`
	EngineLoad       float32   `json:"engineLoad"`
	ElevationAngle   float32   `json:"elevationAngle"`
	O2               float32   `json:"o2"`
	FuelFlow         float32   `json:"fuelFlow"`
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
