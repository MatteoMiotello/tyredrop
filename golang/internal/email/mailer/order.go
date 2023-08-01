package mailer

import (
	"pillowww/titw/internal/email"
	"pillowww/titw/internal/email/templates"
	"pillowww/titw/models"
	"strconv"
)

type OrderMailer struct {
	order     *models.Order
	mailer    *email.Mailer
	templater *email.Templater
}

func NewOrderMailer(order *models.Order) *OrderMailer {
	return &OrderMailer{
		order:     order,
		mailer:    email.NewMailer(),
		templater: email.NewTemplater(),
	}
}

func (r OrderMailer) SendSupportEmail(userFrom *models.User, message string) error {
	subject := "Richiesta di supporto per l'ordine n. #" + strconv.Itoa(int(r.order.ID))

	body, err := r.templater.Process("order_support_email", templates.NewEmailParams(
		subject,
		templates.OrderSupport{
			OrderId:  r.order.ID,
			UserName: userFrom.Email,
			Message:  message,
		}),
	)

	if err != nil {
		return err
	}

	err = r.mailer.SendEmailTo(
		email.From{
			Email: "support@tyresintheworld.com",
			Name:  "Supporto",
		},
		[]string{"dev@pillowww.it"},
		subject,
		body.String(),
	)

	return nil
}
