package domain

//this is the output from training/ data merger app
type TrainingData struct {
	ID                 string `json:"id"`
	Weight             int    `json:"weight"`
	VehicleSpeed       int    `json:"vehicleSpeed"`
	Acceleration       int    `json:"acceleration"`
	WindSpeed          int    `json:"windSpeed"`
	HeadingDirection   int    `json:"headingDirection"`
	WindDirection      int    `json:"windDirection"`
	AtmosphereTemp     int    `json:"atmosphereTemp"`
	AtmospherePressure int    `json:"atmospherePressure"`
	CoolantTemp        int    `json:"coolantTemp"`
	OilPressure        int    `json:"oilPressure"`
	IntakeAirTemp      int    `json:"intakeAirTemp"`
	Humidity           int    `json:"humidity"`
	Rpm                int    `json:"rpm"`
	EngineLoad         int    `json:"engineLoad"`
	ElevationAngle     int    `json:"elevationAngle"`
	O2                 int    `json:"o2"`
	FuelFlow           int    `json:"fuelFlow"`
}

type TrainingDataCollection struct {
	Tdc []TrainingData
}

func (tdc *TrainingDataCollection) AddData(data TrainingData) []TrainingData {
	tdc.Tdc = append(tdc.Tdc, data)
	return tdc.Tdc

}
