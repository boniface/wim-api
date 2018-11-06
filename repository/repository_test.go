package repository

import (
	"testing"
	"time"
	"wim-api/domain"
)

/*func TestTraining(t *testing.T) {

	fmt.Println("Getting data :")
	SelectMerging()
	fmt.Println("The Result is  ")

}*/

func TestOutputCollection(t *testing.T) {
	data1 := domain.OutputData{
		ID:           "sk223",
		Time:         time.Now().Add(67899),
		Year:         time.Now().Year(),
		StartLat:     68,
		StartLon:     95,
		EndLat:       48,
		EndLon:       55,
		Weight:       200,
		IsOverloaded: false,
	}

	data4 := domain.OutputData{
		ID:           "sk423",
		Time:         time.Now().Add(67899),
		Year:         time.Now().Year(),
		StartLat:     68,
		StartLon:     95,
		EndLat:       48,
		EndLon:       55,
		Weight:       200,
		IsOverloaded: false,
	}

	data5 := domain.OutputData{
		ID:           "sk223",
		Time:         time.Now().Add(67899),
		Year:         time.Now().Year(),
		StartLat:     68,
		StartLon:     95,
		EndLat:       48,
		EndLon:       55,
		Weight:       200,
		IsOverloaded: false,
	}

	data2 := domain.OutputData{
		ID:           "sk223",
		Time:         time.Now().Add(3),
		Year:         time.Now().Year(),
		StartLat:     80,
		StartLon:     99,
		EndLat:       82,
		EndLon:       90,
		Weight:       210,
		IsOverloaded: false,
	}

	data3 := domain.OutputData{
		ID:           "sk233",
		Time:         time.Now().Add(1),
		Year:         time.Now().Year(),
		StartLat:     411,
		StartLon:     951,
		EndLat:       482,
		EndLon:       952,
		Weight:       2203,
		IsOverloaded: true,
	}
	data6 := domain.OutputData{
		ID:           "sk234",
		Time:         time.Now().Add(2),
		Year:         time.Now().Year(),
		StartLat:     45,
		StartLon:     44,
		EndLat:       77,
		EndLon:       99,
		Weight:       201,
		IsOverloaded: false,
	}
	d := domain.OutputDataCollection{}
	d.AddData(data1)
	time.Sleep(1000)
	d.AddData(data2)
	time.Sleep(300)
	d.AddData(data3)
	time.Sleep(1000)
	d.AddData(data4)
	d.AddData(data5)
	d.AddData(data6)
	AddOutputDataCollection(d)
	//assert.Equal(t, "200 OK", resp.Status())
}
