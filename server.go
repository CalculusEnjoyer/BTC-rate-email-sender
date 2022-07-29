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
		JSON(context, 200, floatRate)
	} else {
		context.Writer.Header().Set("Content-Type", "text/plain")
		context.String(err.StatusCode, err.Error())
	}
}

func addEmail(context *gin.Context) {
	var email1 Email
	if err := context.Bind(&email1); err != nil {
		return
	} else {
		context.Writer.Header().Set("Content-Type", "text/plain")
		if err1 := addEmailToCsv(email1.Email); err1 == nil {
			context.String(200, "Email has been successfully added")
		} else {
			context.String(err1.StatusCode, err1.Error())
		}
	}
}

func sendAllEmails(context *gin.Context) {
	failedEmails, _ := sendEmails()
	context.Writer.Header().Set("Content-Type", "text/plain")
	joinedEmails := strings.Join(failedEmails, " ")
	if len(failedEmails) == 0 {
		context.String(200, "Emails have been sent")
	} else {
		context.String(200, "Failed to send emails to the following addresses: "+joinedEmails)
	}
}
