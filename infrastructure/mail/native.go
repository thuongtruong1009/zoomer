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

	err := smtp.SendMail(e.cfg.MailHost, auth, e.cfg.MailUser, []string{mail.To}, []byte(mail.Body))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
