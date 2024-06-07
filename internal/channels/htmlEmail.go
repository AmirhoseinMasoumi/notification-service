package channels

import (
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

type HtmlEmailChannel struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewHtmlEmailChannel(host string, port int, username, password string) *HtmlEmailChannel {
	return &HtmlEmailChannel{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (email *HtmlEmailChannel) Send(to string, subject string, htmlBody string) error {
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
