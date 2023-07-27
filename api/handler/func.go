package handler

import (
	"add/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func LoginAuth(auth models.UserLogin, c *gin.Context) (response models.UserInfo, err error) {
	var (
		authResponse models.UserInfo
	)

	responseLogin, err := DoRequest("https://eticket.railway.uz/api/v1/auth/login", "POST", "", auth)
	if err != nil {
		fmt.Println("Error.......1")
		log.Println("error whiling get by id User:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal([]byte(responseLogin), &authResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	return authResponse, nil

}



