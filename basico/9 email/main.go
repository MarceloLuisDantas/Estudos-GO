package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

var sd = []string{
	"smtp.gmail.com",
	"587",
}

func SendEmail(msg []byte, user []string, to []string) string {
	// Authentication.
	auth := smtp.PlainAuth("", user[0], user[1], sd[0])
	// Sending email.
	err := smtp.SendMail(sd[0]+":"+sd[1], auth, user[0], to, msg)
	if err != nil {
		log.Fatal("Erro, ", err)
	}
	return "Email enviado com sucesso, eu acho"
}

func input(label string, s *bufio.Scanner) string {
	fmt.Printf("%s", label)
	s.Scan()
	return s.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	email := input("E-Mail: ", scanner)
	pass := input("Senha: ", scanner)
	envio := input("Destino: ", scanner)
	msg := input("Mensagem: ", scanner)
	SendEmail([]byte(msg), []string{email, pass}, []string{envio})
}
