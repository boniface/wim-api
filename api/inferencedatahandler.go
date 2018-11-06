package api

import (
	"encoding/json"
	"net/http"
	"wim-api/services"
)

func InferenceOutputHandler(w http.ResponseWriter, r *http.Request) {
	res := services.ProcessInference()
	w.Write([]byte("Running Inference \n"))
	formula, _ := json.Marshal(res.Formula)
	w.Write(formula)
	w.Write([]byte("\n"))
	regression, _ := json.Marshal(res)
	w.Write(regression)

	coef, _ := json.Marshal(res.Coef)
	w.Write(coef)

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
