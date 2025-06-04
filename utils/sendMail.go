package utils

import gomail "gopkg.in/gomail.v2"

func SendMail (to, subject, message string) error {
	key := "tcps owxs fyfc iztw"
	from := "ahmed1768476@gmail.com"
	host := "smtp.gmail.com"
	port := 465

	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", message)

	n := gomail.NewDialer(host, port, from, key)

	err := n.DialAndSend(msg)
	
	return err
}