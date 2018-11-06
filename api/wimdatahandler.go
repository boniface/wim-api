package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"wim-api/domain"
	"wim-api/services"
)

func GetOutput(w http.ResponseWriter, r *http.Request) {
	var data domain.OutputDataCollection
	var err error
	//	var y int

	year := chi.URLParam(r, "year")

	ID := chi.URLParam(r, "id")
	if ID != "" {
		if y, err := strconv.ParseUint(year, 10, 32); err == nil {
			data, err = services.ProcessOutputData(ID, int(y))
		}

	}

	if err != nil {
		fmt.Println("Check vehicle ID")
		w.Write([]byte("No data"))
		return
	}

	jdc, err := json.Marshal(data)
	w.Write(jdc)
	/*	for _,od:=range data.Odc{

			jd,err:=json.Marshal(od)
			if err!=nil{
				println("Error while Marshaling ")
				return
			}
			w.Write(jd)

		w.Write([]byte("\t Time:"))
			w.Write([]byte(string(od.Time)))
			w.Write([]byte("\t IsOverloaded: "))
			w.Write([]byte(od.IsOverloaded))

	}*/
}
