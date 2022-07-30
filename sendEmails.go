package main

import (
	"crypto/tls"
	"encoding/csv"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

// created mailbox info with working smtp for quick testing
var senderEmail = "testmyapplication123@outlook.com"
var senderPassword = "genesis2022"
var smtpHost = "smtp.office365.com"
var smtpPort = 587

func sendEmails() ([]string, *RequestError) {
	file, _ := os.Open("emails.csv")
	csvReader := csv.NewReader(file)
	emails, errorRead := csvReader.ReadAll()
	failedEmails := make([]string, 0)
	floatRate, _ := getBtcUahFloat()
	rate := fmt.Sprintf("%v", floatRate)

	message := gomail.NewMessage()
	message.SetHeader("From", senderEmail)
	message.SetHeader("Subject", "BTC/UAH rate")
	message.SetBody("text/plain", "Today's BTC/UAH rate is "+rate)
	domain := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	domain.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if len(emails) == 1 {
		return nil, &RequestError{StatusCode: 500, Err: errors.New("the email database is empty")}
	}

	if errorRead == nil {
		for i := 1; i < len(emails); i++ {
			message.SetHeader("To", emails[i][0])
			if err := domain.DialAndSend(message); err != nil {
				failedEmails = append(failedEmails, emails[i][0])
			}
		}
	}

	if len(failedEmails) == len(emails)-1 {
		return failedEmails, &RequestError{StatusCode: 500, Err: errors.New("all emails have not been sent")}
	} else if len(failedEmails) > 0 {
		return failedEmails, nil
	} else {
		return nil, nil
	}
}
