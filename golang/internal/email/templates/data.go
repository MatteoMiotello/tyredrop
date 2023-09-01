package templates

import (
	"github.com/spf13/viper"
	"pillowww/titw/internal/fs/fshandlers"
)

type EmailParams struct {
	ApplicationUrl  string
	ApplicationLogo string
	Content         any
	Subject         string
}

func NewEmailParams(subject string, content any) *EmailParams {
	fs := fshandlers.NewBrandLogoHandler()

	return &EmailParams{
		ApplicationUrl:  viper.GetString("APPLICATION_URL"),
		ApplicationLogo: fs.GetPublicUrl("logo-transparent.png"),
		Content:         content,
		Subject:         subject,
	}
}
