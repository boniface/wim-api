package domain

//this is the output from training/ data merger app
type TrainingData struct {
	ID                 string  `json:"id"`
	Weight             float32 `json:"weight"`
	VehicleSpeed       float32 `json:"vehicleSpeed"`
	Acceleration       float32 `json:"acceleration"`
	WindSpeed          float32 `json:"windSpeed"`
	HeadingDirection   float32 `json:"headingDirection"`
	WindDirection      float32 `json:"windDirection"`
	AtmosphereTemp     float32 `json:"atmosphereTemp"`
	AtmospherePressure float32 `json:"atmospherePressure"`
	CoolantTemp        float32 `json:"coolantTemp"`
	OilPressure        float32 `json:"oilPressure"`
	IntakeAirTemp      float32 `json:"intakeAirTemp"`
	Humidity           float32 `json:"humidity"`
	Rpm                float32 `json:"rpm"`
	EngineLoad         float32 `json:"engineLoad"`
	ElevationAngle     float32 `json:"elevationAngle"`
	O2                 float32 `json:"o2"`
	FuelFlow           float32 `json:"fuelFlow"`
}

type TrainingDataCollection struct {
	Tdc []TrainingData
}

func (tdc *TrainingDataCollection) AddData(data TrainingData) []TrainingData {
	tdc.Tdc = append(tdc.Tdc, data)
	return tdc.Tdc

}
