package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"math"
	"net/http"
	"wim-api/domain"
	"wim-api/io"
	"wim-api/repository"
	"wim-api/services"
)

func SimpleDataHandler(w http.ResponseWriter, r *http.Request) {

	var data domain.RawInputData

	if err := render.Bind(r, &data); err != nil {
		w.Write([]byte("Error Binding Raw Data"))
		return
	}
	w.Write([]byte("Sending Vehicle data to Process ID :"))
	w.Write([]byte(data.ID))

	vdata, coord := splitData(data)

	w.Write([]byte("Vehicle data"))

	msg, err2 := services.ProcessVehicleData(vdata)
	if err2 != nil {
		fmt.Print("Error occured :", err2)
	}
	fmt.Println("message from svc.procesvehicledata: ", msg)

	w.Write([]byte("\nSending TO Weather API :"))
	owmRes, err1 := io.WeatherAPI(coord)
	if err1 != nil {
		panic(err1)
	}
	services.ProcessWetherData(owmRes)
	w.Write([]byte(owmRes.ID))

}

func BulkVehicleDataHandler(w http.ResponseWriter, r *http.Request) {

	data := []domain.RawInputData{}
	//coordsForAPI:= domain.CoordinateCollection{}
	wdc := domain.WeatherDataCollection{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.Write([]byte("Error on GETBULK DATA"))
		return
	}

	vdc, cc := splitBulkData(data)
	services.ProcessVehicleDataCollection(vdc)

	icoord := cc.Cc[0]
	wdata, err := io.WeatherAPI(icoord)
	if err != nil {
		w.Write([]byte("Error while reading Weather API"))
	}

	for _, coord := range cc.Cc {
		if coordDiff(icoord, coord) { //to avoid frequent calls

			wtemp, err := io.WeatherAPI(coord)
			if err != nil {
				w.Write([]byte("Error while reading  Weather API"))
			}
			icoord = coord
			wdata = wtemp
			wdc.AddData(wdata)
		} else {
			wdata.ID = coord.ID //to insert unique rows
			wdc.AddData(wdata)
		}
	}
	services.ProcessBulkWetherData(wdc)
}

func filterCoords(collection domain.CoordinateCollection) domain.CoordinateCollection {

	ic := collection.Cc[0]
	rc := domain.CoordinateCollection{}
	for _, coord := range collection.Cc {
		if coordDiff(ic, coord) { //to avoid frequent calls
			ic = coord
			rc.AddData(ic)

		} else {
			rc.AddData(ic)

		}
	}
	return rc
}

//if the coordinate difference is higher than the cordinate threshold then call weathermap api
func coordDiff(coordinate1 domain.Coordinate, coordinate2 domain.Coordinate) bool {

	if math.Abs(coordinate1.Latitude-coordinate2.Latitude) > io.Coord_Threshold || math.Abs(coordinate1.Longitude-coordinate2.Longitude) > io.Coord_Threshold {
		return true
	}

	return false
}

func ApiKeySetter() string {
	return repository.SelectKey()
}
