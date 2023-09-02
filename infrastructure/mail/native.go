package mail

import (
	"fmt"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"net/smtp"
	"strings"
)

type mail struct {
	cfg *configs.Configuration
}

func NewMail(cfg *configs.Configuration) IMail {
	return &mail{
		cfg: cfg,
	}
}

func (m *mail) SendingMail(mail *Mail) error {
	auth := smtp.PlainAuth("", m.cfg.MailUser, m.cfg.MailPassword, strings.Split(m.cfg.MailHost, ":")[0])

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\nThis is auto message from Zoomer\n\n%s", mail.To, mail.Subject, mail.Body))

	err := smtp.SendMail(m.cfg.MailHost, auth, m.cfg.MailUser, []string{mail.To}, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
