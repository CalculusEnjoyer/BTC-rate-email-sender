package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/rate", getBtcUah)
	router.POST("/subscribe", addEmail)
	router.POST("/sendEmails", sendAllEmails)
	router.Run() //runs on localhost:8080 by default
}
