package mailer

import (
	"github.com/friendsofgo/errors"
	"github.com/spf13/viper"
	"pillowww/titw/internal/email"
	"pillowww/titw/internal/email/templates"
	"pillowww/titw/models"
)

type ResetPasswordMailer struct {
	ResetPassword *models.ResetPassword
	mailer        *email.Mailer
	templater     *email.Templater
}

type ResetPassword struct {
	Url string
}

func NewResetPasswordMailer(rp *models.ResetPassword) *ResetPasswordMailer {
	return &ResetPasswordMailer{
		ResetPassword: rp,
		mailer:        email.NewMailer(),
		templater:     email.NewTemplater(),
	}
}

func (r ResetPasswordMailer) SendResetEmail(u *models.User) error {
	subject := "Richiesta reset password"

	body, err := r.templater.Process("reset_password_email", templates.NewEmailParams(
		subject,
		ResetPassword{
			Url: viper.GetString("APPLICATION_FRONTEND_URL") + "/reset_password/" + r.ResetPassword.Token,
		},
	))

	if err != nil {
		return errors.WithMessage(err, "error generating template")
	}

	err = r.mailer.SendEmailTo(
		email.From{
			Email: "noreply@tyresintheworld.com",
			Name:  "NoReply - Tyres in the world",
		},
		[]string{
			u.Email,
		},
		subject,
		body.String(),
	)

	return nil
}
