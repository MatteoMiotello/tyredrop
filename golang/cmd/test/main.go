package main

import (
	"context"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/email/mailer"
	"pillowww/titw/models"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
}

func main() {
	ctx := context.Background()

	one, err := models.Orders().One(ctx, db.DB)

	if err != nil {
		panic(err)
	}

	billing, err := one.UserBilling().One(ctx, db.DB)

	if err != nil {
		panic(err)
	}

	u, err := billing.User().One(ctx, db.DB)

	if err != nil {
		panic(err)
	}

	m := mailer.NewOrderMailer(one)

	err = m.SendSupportEmail(u, "Prima email")

	if err != nil {
		panic(err)
	}
}
