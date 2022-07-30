# BTC-rate-email-sender
This is an application that provides an opportunity to get live BTC/UAH rate and 
send it to subscribed emails.

## Installation
Open terminal in the project directory and run the following: 
```sh
docker build -t btc-rate .
docker run -p 8080:80 btc-rate
```
Application will be running on localhost:8080.

## Application structure
Gomail package was used for sending emails
https://github.com/go-gomail/gomail <br />
Gin framework was used for building app https://github.com/gin-gonic/gin <br />
<br />
All user data are stored in email.csv file with one column called email. emailDb.go is 
responsible for adding new emails to the file.<br />
<br />
sendEmail.go sends emails to subscribed users. It contains global variables with the sender email information 
(it already has mailbox info that is created for quick testing).<br />
<br />
btcRate.go contains function for getting up-to-date BTC/UAH rate using kuna.io API <br />
<br />
error.go has RequestError struct that all functions return, so it makes easier to deal with status codes and reasons of errors for server.go.<br />

## Routes
