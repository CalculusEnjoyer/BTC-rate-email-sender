package main

import (
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

var senderEmail = "kravchuk883@gmail.com"
var senderPassword = "hoplay009"
var smtpHost = "smtp.office365.com"
var smtpPort = 587

func sendEmails() ([]string, error) {
	file, _ := os.Open("emails.csv")
	csvReader := csv.NewReader(file)
	emails, errorRead := csvReader.ReadAll()
	failedEmails := make([]string, 0)
	floatRate, _ := getBtcUahFloat()
	rate := fmt.Sprintf("%v", floatRate)

	message := gomail.NewMessage()
	message.SetHeader("From", "kravchuk883@gmail.com")
	message.SetHeader("Subject", "BTC/UAH rate")
	message.SetBody("text/plain", "Today's BTC/UAH rate is "+rate)
	domain := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	domain.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if errorRead == nil {
		for i := 1; i < len(emails); i++ {
			message.SetHeader("To", emails[i][0])
			if err := domain.DialAndSend(message); err != nil {
				failedEmails = append(failedEmails, emails[i][0])
			}
		}
	}
	return failedEmails, nil
}
