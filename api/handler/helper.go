package handler

import (
	"fmt"
	"net/smtp"
	"time"
)
// data to'grlash uchun
func FormatDate(dateString string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return "", err
	}
	formattedDate := parsedDate.Format("02.01.2006")
	return formattedDate, nil
}

//gmailga sms yuborish
func SendMail(sms string) {
	auth := smtp.PlainAuth(
		"",
		"onedevaloper@gmail.com",
		"ubmcamzbygeuopiv",
		"smtp.gmail.com",
	)

	senMsg := "Notification from the server\n" + sms

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"madboevshuhrat@gmail.com",
		[]string{"madboevshuhrat@gmail.com"},
		[]byte(senMsg),
	)

	if err != nil {
		fmt.Println(err)
	}
}