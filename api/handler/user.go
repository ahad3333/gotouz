package handler

import (
	"add/config"
	"add/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestBody godoc
// @ID RequestBody
// @Router /user [POST]
// @Summary Request
// @Description Request
// @Tags Request
// @Accept json
// @Produce json
// @Param book body models.RequestBody true "RequestBody"
// @Success 200 {object} models.ResponseForFront "GetBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func RequestBody(c *gin.Context) {
	var (
		req           models.RequestBody
		reqdirection  []models.Direction
		responseLogin models.UserInfo
		checkUser int
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("req....", req)
	if req.StationFrom == req.StationTo {
		log.Println("Choose 2 different Stations:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Choose 2 different Stations",
		})
		return
	}

	for {
		fmt.Println("test......1",time.Now())
		if checkUser > 4 {
			checkUser = 0
		}
		if checkUser > 0 {
			requestLogin:= Check(checkUser)
			responseLogin, err = LoginAuth(requestLogin, c)
		}else{
			requestLogin := models.UserLogin{
				Username: config.LoginGmail4,
				Password: config.PasswordGmail4,
			}
			responseLogin, err = LoginAuth(requestLogin, c)
		}


			if responseLogin.Token == "" {
				fmt.Println("tokkin yo'q", responseLogin.Token)
			}

		for _, val := range req.Direction {

			res, err := FormatDate(val.DepDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			direction := models.Direction{
				DepDate: res,
				Fullday: val.Fullday,
				Type:    val.Type,
			}
			reqdirection = append(reqdirection, direction)
		}
		fmt.Println("• Looking for a ticket for you.......")
		responsekey, err := Request(models.RequestBody{
			Direction:         reqdirection,
			StationFrom:       req.StationFrom,
			StationTo:         req.StationTo,
			DetailNumPlaces:   req.DetailNumPlaces,
			ShowWithoutPlaces: req.ShowWithoutPlaces,
			Wagon:             req.Wagon,
			Time:              req.Time,
			Types:             req.Types,
		}, responseLogin, c)

		if err != nil {
			fmt.Println("Error.......2")
		}

		fmt.Println("responsekey...",responsekey)

		reqUser := models.User{
			UserID:      "04f428bd-4a46-460b-bc6c-bb2552d79679",
			Firstname:   "test uchun",
			Midname:     "test uchun",
			Lastname:    "test Uchun",
			Sex:         "M",
			BirthDay:    "09.08.1990",
			Citizenship: "UKR",
			DocType:     "33",
			RegionID:    "",
			Doc:         "63526",
			YourSelf:    false,
		}

		payload := `{"backward": false, "ordersRequest": [{"stationFrom": ` + `"` + req.StationFrom + `"` + `, "stationTo": ` + `"` + req.StationTo + `"` + `, "depDate": ` + `"` + responsekey.Depdate + `"` + `, "train": {"number": ` + `"` + responsekey.Number + `"` + `}, "car": {"type": ` + `"` + responsekey.TypeShow + `"` + `, "number": ` + `"` + responsekey.Respkey + `"` + `}, "requirements": {"seatsRange": ` + `"` + responsekey.Response + `"` + `}}], "withInsurance": false, "withSmsNotification": false, "phone": "998", "passengers": [{"children": null, "valid": true, "child5From10": false, "birthDay": ` + `"` + reqUser.BirthDay + `"` + `, "citizenship": ` + `"` + reqUser.Citizenship + `"` + `, "doc": "63526", "docType": "ЗЗ", "firstname": ` + `"` + reqUser.Firstname + `"` + `, "lastname": ` + `"` + reqUser.Lastname + `"` + `, "midname": ` + `"` + reqUser.Midname + `"` + `, "regionId": "", "pinfl": "", "sex": ` + `"` + reqUser.Sex + `"` + `}]}`

		respOrder := DoRequest2("Bearer "+responseLogin.Token, payload)

		fmt.Println("OrderId------>",respOrder)
		if respOrder == config.Error {
			fmt.Println("error....")
			checkUser++
		}

		current := time.Now()
		timeS := current.Format("15:04:05")

		resp := models.ResponseForFront{
			Brand:   " 🚄 "+"Afrosiyob",
			Places:  " 📍 "+responsekey.Response,
			Carnum:  " 📍 "+responsekey.Respkey,
			Date:    date + " 🕖 • " + responsekey.Depdate +" 📆 ",
			OrderId: respOrder,
			NowTime: " 🕖 " +timeS,
			Token:   responseLogin.Token,
		}

		parsedUUID, err := uuid.Parse(respOrder)

		if err != nil {
			fmt.Println("Invalid UUID:", err)
		}

		if `"`+parsedUUID.String()+`"` == respOrder {
			responseGmail := "• OrderId:\t" + string(respOrder) + "\n" + "• Brand:\tAfrosiyob\n" + "• Places:\t" + responsekey.Response + "\n" + "• Carnum:\t" + responsekey.Respkey + "\n" + "• Date:\t" + date + "\n" + "Now time:\t" + timeS
			SendMail(responseGmail)
			c.JSON(200, resp)
			break
		}
		fmt.Println("test......10",time.Now())

	}
}
