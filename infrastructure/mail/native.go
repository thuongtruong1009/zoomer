package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
)

type mail struct {
	cfg *configs.Configuration
}

func NewMail(cfg *configs.Configuration) IMail {
	return &mail{
		cfg: cfg,
	}
}

func (e *mail) SendingNativeMail(mail *Mail) error {
	auth := smtp.PlainAuth("", e.cfg.MailUser, e.cfg.MailPassword, strings.Split(e.cfg.MailHost, ":")[0])

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\nThis is auto message from Zoomer\n\n%s", mail.To, mail.Subject, mail.Body))
	err := smtp.SendMail(e.cfg.MailHost, auth, e.cfg.MailUser, []string{mail.To}, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
