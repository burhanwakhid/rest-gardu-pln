package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"burhanwakhid.space/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type GarduController struct{}

type result struct {
	Area       string
	Distribusi string
	JmlPlgn    string
	Kva        string
	Coverage   string
	Distance   string
}

type GarduInput struct {
	Lat    string `json:"lat" binding:"required"`
	Long   string `json:"long" binding:"required"`
	Radius string `json:"radius" binding:"required"`
}

func (ctrl GarduController) ListGardu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput GarduInput

	if err := c.ShouldBind(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error})
	}

	lat, errLat := strconv.ParseFloat(dataInput.Lat, 10)
	if errLat != nil {
		// insert error handling here
	}

	long, errLong := strconv.ParseFloat(dataInput.Long, 10)
	if errLong != nil {
		// insert error handling here
	}

	radius, err := strconv.Atoi(dataInput.Radius)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	// lat := -7.785973
	// long := 110.399957
	// var gardu []model.GarduM
	var res []result

	if err := db.Raw(`SELECT area, distribusi , jml_plgn ,kva ,coverage , 
		(6371e3 * acos (cos ( radians(?) ) 
		* cos( radians( x ) )
		* cos( radians( y ) - radians(?) )
		+ sin ( radians(?) )* sin( radians( x ) ))) 
		AS distance 
		FROM gardu_ms 
		HAVING distance <= ?`, lat, long, lat, radius).
		Scan(&res); err != nil {
		c.JSON(http.StatusOK, gin.H{"data": res})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}

}

func (ctrl GarduController) AddListGardu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	url := "http://apps.iconpln.co.id:7181/Panas-1.0/AssetGarduAllRf"

	spaceClient := http.Client{
		Timeout: time.Second * 400, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	listGardu := model.JsonGardu{}

	jsonErr := json.Unmarshal(body, &listGardu)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// data := model.GarduM{}
	// db.Create(&model.GarduM{
	// 	Kva:     "sfad",
	// 	JmlPlgn: "sfad",
	// })

	for _, v := range listGardu.ListAsset {

		lat, errLat := strconv.ParseFloat(v.X, 10)
		if errLat != nil {
			// insert error handling here
		}

		long, errLong := strconv.ParseFloat(v.Y, 10)
		if errLong != nil {
			// insert error handling here
		}

		db.Create(&model.GarduM{
			Kva:          v.Kva,
			JmlPlgn:      v.JmlPlgn,
			StatusNyala:  v.StatusNyala,
			Distribusi:   v.Distribusi,
			Area:         v.Area,
			Unit:         v.Unit,
			NoTiang:      v.NoTiang,
			JenisTrafo:   v.JenisTrafo,
			X:            lat,
			Y:            long,
			Coverage:     v.Coverage,
			KodeLaporan:  v.KodeLaporan,
			KodeLaporanP: v.KodeLaporanP,
			KodeLaporanJ: v.KodeLaporanJ,
		})
		fmt.Println(v.Area)
	}

}
