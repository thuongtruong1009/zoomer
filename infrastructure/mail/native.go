package mail

import (
	"bytes"
	"fmt"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"html/template"
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

func (e *mail) SendingNativeMail(mail *Mail) error {
	auth := smtp.PlainAuth("", e.cfg.MailUser, e.cfg.MailPassword, strings.Split(e.cfg.MailHost, ":")[0])

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", mail.Subject, "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    mail.To,
		Message: mail.Body,
	})

	err := smtp.SendMail(e.cfg.MailHost, auth, e.cfg.MailUser, []string{mail.To}, body.Bytes())

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
