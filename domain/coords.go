package domain

import "time"

type Coordinate struct {
	ID        string    `json: "id"`
	Time      time.Time `json:"time"`
	Latitude  float32   `json:"lat"`
	Longitude float32   `json:"lon"`
}

type CoordinateCollection struct {
	Cc []Coordinate
}

func (cc *CoordinateCollection) AddData(data Coordinate) []Coordinate {
	cc.Cc = append(cc.Cc, data)
	return cc.Cc

}
