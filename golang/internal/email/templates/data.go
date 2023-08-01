package templates

import "github.com/spf13/viper"

type EmailParams struct {
	ApplicationUrl string
	Content        any
	Subject        string
}

func NewEmailParams(subject string, content any) *EmailParams {
	return &EmailParams{
		ApplicationUrl: viper.GetString("APPLICATION_URL"),
		Content:        content,
		Subject:        subject,
	}
}

type OrderSupport struct {
	EmailParams
	UserName string
	OrderId  int64
	Message  string
}
