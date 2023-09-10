package mailer

import (
	"github.com/spf13/viper"
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

type OrderSupportData struct {
	UserEmail   string
	OrderNumber string
	OrderDate   string
	OrderUrl    string
	Message     string
}

func NewOrderMailer(order *models.Order) *OrderMailer {
	return &OrderMailer{
		order:     order,
		mailer:    email.NewMailer(),
		templater: email.NewTemplater(),
	}
}

func (r OrderMailer) SendSupportEmail(userFrom *models.User, message string) error {
	subject := "Richiesta di supporto per l'ordine n. #" + r.order.OrderNumber.String

	params := templates.NewEmailParams(
		subject,
		OrderSupportData{
			OrderNumber: r.order.OrderNumber.String,
			UserEmail:   userFrom.Email,
			Message:     message,
			OrderUrl:    viper.GetString("APPLICATION_FRONTEND_URL") + "/admin/order/" + strconv.Itoa(int(r.order.ID)),
			OrderDate:   r.order.CreatedAt.Format("02/01/2006"),
		})

	body, err := r.templater.Process("order_support_email", params)

	if err != nil {
		return err
	}

	b := body.String()

	err = r.mailer.SendEmailTo(
		email.NewNoreplyFrom(),
		[]string{viper.GetString("SUPPORT_EMAIL")},
		subject,
		b,
	)

	return nil
}
