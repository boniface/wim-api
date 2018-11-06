package ml

import (
	"fmt"
	"github.com/sivaramalingamk/regression"
	"wim-api/domain"
)

func Regression(collection domain.TrainingDataCollection) *regression.Regression {

	r := new(regression.Regression)

	r.SetObserved("Weight")
	r.SetVar(0, "Acceleration")
	r.SetVar(1, "Vehicle_Speed")
	r.SetVar(2, "Heading_Direction")
	r.SetVar(3, "Elevation_Angle")
	r.SetVar(4, "Engine_Load")
	r.SetVar(5, "Coolent_Temp")
	r.SetVar(6, "Oil_Pressure")

	r.SetVar(7, "O2_Flow")
	r.SetVar(8, "Fuel_Flow")

	r.SetVar(9, "RPM")
	r.SetVar(10, "Intake_Air_Temp")
	r.SetVar(11, "Atmosphere_Temp")
	//r.SetVar(12, "Atmosphere_Pressure")
	//r.SetVar(13, "Humidity")
	//r.SetVar(14, "Wind_Speed")
	//r.SetVar(15, "Wind_Direction")

	for _, td := range collection.Tdc {
		r.Train(regression.DataPoint(float64(td.Weight), []float64{float64(td.Acceleration), float64(td.VehicleSpeed), float64(td.HeadingDirection), float64(td.ElevationAngle), float64(td.EngineLoad), float64(td.CoolantTemp), float64(td.OilPressure), float64(td.O2), float64(td.FuelFlow), float64(td.Rpm), float64(td.IntakeAirTemp), float64(td.AtmosphereTemp) /*,float64(td.AtmospherePressure),float64(td.Humidity),float64(td.WindSpeed),float64(td.WindDirection)*/}))
	}

	r.AddCross(regression.MultiplierCross(1, 2))
	r.AddCross(regression.MultiplierCross(1, 3))
	r.AddCross(regression.MultiplierCross(2, 3, 4))
	r.AddCross(regression.MultiplierCross(3, 4))
	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)

	for i, c := range r.Coef {
		if i == 0 {

			continue
		}
		if c < 0 {
			fmt.Println("Coefficient is negative, but shouldn't be: %.2f", c)
		}
	}

	if r.R2 < 0.8 {
		fmt.Println("R^2 was %.2f, but we expected > 80", r.R2)
	}

	prediction, err := r.Predict([]float64{1.01, 15.50, 30.00, 19.43, 4.39, 50.00, 1220.00, 34.00, 24.00, 999.00, 233.00, 303.20})
	fmt.Println("Predicted weight is :", prediction, " With error: ", err)

	return r
}
