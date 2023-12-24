package lib

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(err string) {
	from := os.Getenv("TinoBackUpEmailAccount")
	frompassword := os.Getenv("TinoBackUpEmailPassword")
	to := "konstantin.v.milchev@gmail.com"

	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	auth := smtp.PlainAuth("", from, frompassword, smtpServer)
	serverAddress := fmt.Sprintf("%s:%d", smtpServer, smtpPort)

	smtp.SendMail(serverAddress, auth, from, []string{to}, []byte(err))
}
