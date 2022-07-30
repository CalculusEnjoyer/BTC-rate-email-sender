# BTC-rate-email-sender
This is an application that provides an opportunity to get live BTC/UAH rate and 
send it to subscribed emails.

## Installaiton
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
All user data are stored in email.csv file. It contains one column called email. EmailDb.go is 
responsible for adding new emails to the file.
<br />
