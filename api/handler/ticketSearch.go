package handler

import (
	"add/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var date string

func Request(req models.RequestBody, auth models.UserInfo, c *gin.Context) (models.ResponseKey, error) {
	var (
		respo              models.ResponseAll
		directionRespArrey []models.Direction2
	)

	requestSearch := models.RequestSearch{
		Direction:         req.Direction,
		StationFrom:       req.StationFrom,
		StationTo:         req.StationTo,
		DetailNumPlaces:   req.DetailNumPlaces,
		ShowWithoutPlaces: req.ShowWithoutPlaces,
	}

	fmt.Println("Req.....time...", req.Time)

	for {

		res, err := DoRequest("https://eticket.railway.uz/api/v1/trains/availability/space/between/stations", "POST", auth.Token, requestSearch)

		if err != nil {
			fmt.Println("Error.......1")
		}

		err = json.Unmarshal([]byte(res), &respo)
		if err != nil {
			fmt.Println("Error:", err)
			return models.ResponseKey{}, err
		}

		respGmail := ""
		for _, direction := range respo.Express.Direction {
			for _, trainData := range direction.Trains {
				for _, train := range trainData.Train {
					if train.Brand == "Afrasiyab" {
						if train.Departure.Time == req.Time {
							for _, a := range train.Places.Cars {
								currentTime := time.Now()
								timeString := currentTime.Format("15:04:05")
								respGmail = "• Brand:\t" + train.Brand + "\n" + "• SeatsUndef:\t" + a.Seats.SeatsUndef + "\n" + "• Date:\t" + train.Departure.Time + "\t" + train.Arrival.LocalDate + "\n" + "Now time:\t" + timeString
								date = train.Departure.Time
								fmt.Println(respGmail)
								fmt.Println("----------------------------")
								directionResp := models.Direction2{
									DepDate: train.Arrival.LocalDate,
									Fullday: true,
									Type:    direction.Type,
									Train: models.Train{
										Number: train.Number,
									},
								}
								directionRespArrey = append(directionRespArrey, directionResp)
								resp := models.TrainAddresses{
									StationFrom:     req.StationFrom,
									StationTo:       req.StationTo,
									DetailNumPlaces: req.DetailNumPlaces,
									Direction:       directionRespArrey,
								}
								reqVagon, err := Vagon(resp, auth, c)

								if err != nil {
									fmt.Println("Error.......3")
								}
								var (
									depdate  string
									number   string
									typeShow string
								)

								for _, val := range resp.Direction {
									depdate = val.DepDate
									number = val.Train.Number
								}
								placesMap := make(map[string][]string)
								for _, val3 := range reqVagon.Direction {
									for _, t := range val3.Trains {
										for _, r := range t.Train.Cars {
											typeShow = r.TypeShow
											for _, c := range r.Car {
												placesMap[c.TypeOfCarSchema.CarNumber] = append(placesMap[c.TypeOfCarSchema.CarNumber], c.Places)
											}
										}
									}
								}
								if req.Types == "xxx" {
									response, respkey := Preparation2(placesMap)
									responseKey := models.ResponseKey{
										Response: response,
										Respkey:  respkey,
										Depdate:  depdate,
										Number:   number,
										TypeShow: typeShow,
									}
									if response != "" && respkey != "" {
										return responseKey, nil
									}
								} else {
									response, respkey := Preparation(placesMap, req.Types)
									responseKey := models.ResponseKey{
										Response: response,
										Respkey:  respkey,
										Depdate:  depdate,
										Number:   number,
										TypeShow: typeShow,
									}

									if response != "" && respkey != "" {
										return responseKey, nil
									}
								}
							}
						}

					}
				}
			}
		}

		// time.Sleep(500 * time.Millisecond) // 0.5 Second
		time.Sleep(1 * time.Second) // 1 Second
		// time.Sleep(2 * time.Second) // 2 Second
	}
}

func Vagon(req models.TrainAddresses, auth models.UserInfo, c *gin.Context) (resp models.ResponseVagon, err error) {

	responsevagon, err := DoRequest("https://eticket.railway.uz/api/v2/trains/given/availability/place", "POST", "Bearer "+auth.Token, req)

	if err != nil {
		fmt.Println("Error.......3")
	}

	var requestVagon models.ResponseVagon

	err = json.Unmarshal([]byte(responsevagon), &requestVagon)
	if err != nil {
		fmt.Println("Error:...requestVagon", err)
		return
	}

	resp = requestVagon

	return resp, nil
}
