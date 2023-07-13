package mail

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type SmtpMail interface {
	SendVerificationEmail(toEmail string, code string, subject string)
	SendEmailWelcome(toEmail string, subject string)
}

type SmtpMailImpl struct{}

// SendVerificationEmail implements SmtpMail.
func (*SmtpMailImpl) SendVerificationEmail(toEmail string, code string, subject string) {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("SMTP_MAIL_SENDER"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", "Hello <b>"+toEmail+"</b>! Welcome to Gobook! Your verification code is <b>"+code+"</b>")

	port, portERR := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if portERR != nil {
		panic(portERR)
	}

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func (mi *SmtpMailImpl) SendEmailWelcome(toEmail string, subject string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "wafiyanwarulhikam12@gmail.com")
	m.SetHeader("To", "andikapratama5689@gmail.com")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", "Hello <b>"+toEmail+"</b>! Welcome to GoBook!")

	port, portERR := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if portERR != nil {
		panic(portERR)
	}

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, "400aa71a30e23c", os.Getenv("SMTP_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func NewSmtpMail() SmtpMail {
	return &SmtpMailImpl{}
}
