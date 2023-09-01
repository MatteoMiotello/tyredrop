package mailer

import (
	"github.com/friendsofgo/errors"
	"github.com/spf13/viper"
	"pillowww/titw/internal/email"
	"pillowww/titw/internal/email/templates"
)

type SupportEmailData struct {
	Email   string
	Name    string
	Phone   string
	Message string
}

type SupportMailer struct {
	mailer    *email.Mailer
	templater *email.Templater
}

func NewSupportMailer() *SupportMailer {
	return &SupportMailer{
		mailer:    email.NewMailer(),
		templater: email.NewTemplater(),
	}
}

func (r SupportMailer) SendResetEmail(emailAddr string, phone string, name string, message string) error {
	subject := "Richiesta di Assistenza da un Cliente"

	body, err := r.templater.Process("support_email", templates.NewEmailParams(
		subject,
		SupportEmailData{
			Email:   emailAddr,
			Phone:   phone,
			Name:    name,
			Message: message,
		},
	))

	if err != nil {
		return errors.WithMessage(err, "error generating template")
	}

	err = r.mailer.SendEmailTo(
		email.NewNoreplyFrom(),
		[]string{
			viper.GetString("SUPPORT_EMAIL"),
		},
		subject,
		body.String(),
	)

	if err != nil {
		return err
	}

	return nil
}
