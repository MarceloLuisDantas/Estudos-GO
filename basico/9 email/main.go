package main

import (
	"log"
	"net/smtp"
)

var user = []string{
	"dromedariodechapeu@gmail.com",
	"1254marcelo",
}

var sd = []string{
	"smtp.gmail.com",
	"587",
}

func SendEmail(msg []byte) string {
	to := []string{
		"marceloluis257@gmail.com",
	}
	// Authentication.
	auth := smtp.PlainAuth("", user[0], user[1], sd[0])
	// Sending email.
	err := smtp.SendMail(sd[0]+":"+sd[1], auth, user[9], to, msg)
	if err != nil {
		log.Fatal("Erro, ", err)
	}
	return "Email enviado com sucesso, eu acho"
}

func main() {
	msg := []byte("Gostoso")
	SendEmail(msg)
}
