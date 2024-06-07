package channels

import (
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

type EmailChannel struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailChannel(host string, port int, username, password string) *EmailChannel {
	return &EmailChannel{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (email *EmailChannel) Send(to string, subject string, body string) error {
	msg := "From: " + email.Username + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(fmt.Sprintf("%s:%d", email.Host, email.Port),
		smtp.PlainAuth("", email.Username, email.Password, email.Host),
		email.Username, []string{to}, []byte(msg))

	if err != nil {
		log.Info().Msgf("smtp error: %s", err)
		return err
	}
	log.Info().Msgf("Successfully sent email to " + to)
	return nil
}

func (email *EmailChannel) SendHTML(to string, subject string, htmlBody string) error {
	msg := "From: " + email.Username + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		htmlBody

	err := smtp.SendMail(fmt.Sprintf("%s:%d", email.Host, email.Port),
		smtp.PlainAuth("", email.Username, email.Password, email.Host),
		email.Username, []string{to}, []byte(msg))

	if err != nil {
		log.Info().Msgf("smtp error: %s", err)
		return err
	}
	log.Info().Msgf("Successfully sent HTML email to %s", to)
	return nil
}
