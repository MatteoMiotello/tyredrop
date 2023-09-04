package mailer

import (
	"github.com/friendsofgo/errors"
	"github.com/spf13/viper"
	"pillowww/titw/internal/email"
	"pillowww/titw/internal/email/templates"
	"pillowww/titw/models"
	"strconv"
)

type UserMailer struct {
	user      *models.User
	mailer    *email.Mailer
	templater *email.Templater
}

type NewUserData struct {
	UserName         string
	UserEmail        string
	UserUrl          string
	SubscriptionDate string
}

func NewUserMailer(u *models.User) *UserMailer {
	return &UserMailer{
		user:      u,
		mailer:    email.NewMailer(),
		templater: email.NewTemplater(),
	}
}

func (r UserMailer) SendNewUserNotification() error {
	subject := "Nuova iscrizione utente"

	body, err := r.templater.Process("new_user_notification_email", templates.NewEmailParams(
		subject,
		NewUserData{
			UserEmail:        r.user.Email,
			UserName:         r.user.Name + " " + r.user.Surname.String,
			UserUrl:          viper.GetString("APPLICATION_FRONTEND_URL") + "/admin/user/" + strconv.Itoa(int(r.user.ID)),
			SubscriptionDate: r.user.CreatedAt.Format("02/01/2006"),
		},
	))

	if err != nil {
		return errors.WithMessage(err, "error generating template")
	}

	err = r.mailer.SendEmailTo(
		email.NewNoreplyFrom(),
		[]string{
			viper.GetString("NOTIFICATION_EMAIL"),
		},
		subject,
		body.String(),
	)

	if err != nil {
		return err
	}

	return nil
}
