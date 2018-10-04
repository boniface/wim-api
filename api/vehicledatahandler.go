package api

import (
	"fmt"
	"net/http"
	"wim-api/domain"
)

func VehicleDataHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vehicledata, ok := ctx.Value("id").(*domain.VehicleData)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("Vechicle Id is:%s", vehicledata.ID)))
}

func BulkVehicleDataHandler() {

}
