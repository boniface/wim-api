package domain

type Coordinate struct {
	ID        string  `json: "id"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type CoordinateCollection struct {
	Cc []Coordinate
}

func (cc *CoordinateCollection) AddData(data Coordinate) []Coordinate {
	cc.Cc = append(cc.Cc, data)
	return cc.Cc

}
