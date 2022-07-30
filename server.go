package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Email struct {
	Email string `json:"email" form:"email"`
}

func JSON(c *gin.Context, code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, obj)
}

func getBtcUah(context *gin.Context) {
	if floatRate, err := getBtcUahFloat(); err == nil {
		JSON(context, 200, int(floatRate))
	} else {
		JSON(context, err.StatusCode, err.Error())
	}
}

func addEmail(context *gin.Context) {
	var email1 Email
	if err := context.Bind(&email1); err != nil {
		return
	} else {
		if err1 := addEmailToCsv(email1.Email); err1 == nil {
			JSON(context, 200, "Email has been successfully added")
		} else {
			JSON(context, err1.StatusCode, err1.Error())
		}
	}
}

func sendAllEmails(context *gin.Context) {
	failedEmails, err := sendEmails()
	joinedEmails := strings.Join(failedEmails, " ")
	if failedEmails == nil && err == nil {
		JSON(context, 200, "Emails have been sent")
	} else if err == nil {
		JSON(context, 200, "Failed to send emails to the following addresses: "+joinedEmails)
	} else {
		JSON(context, err.StatusCode, err.Error())
	}
}
