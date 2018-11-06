package repository

import (
	"fmt"
	"wim-api/domain"
)

func SelectOutputData(id string, year int) (domain.OutputDataCollection, error) {

	fmt.Println("Selecting OUTPUT data From Database with ID: ", id, " year:", year)
	defer getCluster().Close()
	odc := domain.OutputDataCollection{}
	temp := domain.OutputData{}
	iter := Session.Query(`SELECT id,time,year,startlat,startlon,endlat,endlon,isoverloaded,weight FROM outputdata WHERE id=? and year=?`, id, year).Iter()
	for iter.Scan(&temp.ID, &temp.Time, &temp.Year, &temp.StartLat, &temp.StartLon, &temp.EndLat, &temp.EndLon, &temp.IsOverloaded, &temp.Weight) {
		odc.AddData(temp)
	}
	if err := iter.Close(); err != nil {
		fmt.Println("Error @ selecting output data ", err)
		return odc, err
	}
	fmt.Println("Selected ", len(odc.Odc), " Number of Output Data")
	return odc, nil
}

func AddOutputDataCollection(odc domain.OutputDataCollection) (string, error) {
	fmt.Println(" **** Creating new Output data ****\n Inserting ", len(odc.Odc), "number of rows")
	defer getCluster().Close()
	for _, data := range odc.Odc {

		if err := Session.Query("INSERT INTO outputdata(id,time,year,startlat,startlon,endlat,endlon,weight,isoverloaded) values(?,?,?,?,?,?,?,?,?)", data.ID, data.Time, data.Year, data.StartLat, data.StartLon, data.EndLat, data.EndLon, data.Weight, data.IsOverloaded).Exec(); err != nil {
			fmt.Println("Error while inserting Output Data", err)
			fmt.Println(err)
			return "", err
		}

	}
	return "Success", nil
}
