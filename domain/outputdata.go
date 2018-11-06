package domain

import "time"

type OutputData struct {
	ID           string    `json:"id"`
	Year         int       `json:"year"`
	Time         time.Time `json:"time"`
	StartLat     float32   `json:"startLat"`
	StartLon     float32   `json:"startLon"`
	EndLat       float32   `json:"EndLat"`
	EndLon       float32   `json:"EndLon"`
	Weight       float32   `json:"weight"`
	IsOverloaded bool      `json:"isOverloaded"`
}

type OutputDataCollection struct {
	Odc []OutputData
}

func (odc *OutputDataCollection) AddData(data OutputData) []OutputData {
	odc.Odc = append(odc.Odc, data)
	return odc.Odc

}
