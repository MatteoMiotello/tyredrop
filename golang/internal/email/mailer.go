package email

import (
	"fmt"
	"github.com/resendlabs/resend-go"
	"github.com/spf13/viper"
)

type Mailer struct {
	client    *resend.Client
	templater *Templater
}

func NewMailer() *Mailer {
	fmt.Println(viper.GetString("RESEND_API_TOKEN"))

	return &Mailer{
		resend.NewClient(viper.GetString("RESEND_API_TOKEN")),
		NewTemplater(),
	}
}

func (m *Mailer) SendEmail(request *resend.SendEmailRequest) error {
	_, err := m.client.Emails.Send(request)

	if err != nil {
		return err
	}

	return nil
}

func (m *Mailer) SendEmailTo(from From, to []string, subject string, body string) error {
	req := &resend.SendEmailRequest{
		From:    from.String(),
		To:      to,
		Subject: subject,
		Html:    body,
	}

	return m.SendEmail(req)
}
