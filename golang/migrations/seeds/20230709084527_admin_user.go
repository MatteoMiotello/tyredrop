package seeds

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
	"pillowww/titw/pkg/security"
)

func init() {
	goose.AddMigration(upAdminUser, downAdminUser)
}

func upAdminUser(tx *sql.Tx) error {
	ctx := context.Background()
	userDao := user.NewDao(tx)
	defPass := viper.GetString("DEFAULT_PASSWORD")
	hashed, err := security.HashPassword(defPass)

	if err != nil {
		return err
	}

	role, err := userDao.FindUserRoleByCode(ctx, user.ADMIN_ROLE)

	if err != nil {
		return err
	}

	newUser := &models.User{
		UserRoleID:        role.ID,
		DefaultLanguageID: 1,
		Email:             viper.GetString("DEFAULT_EMAIL"),
		Username:          null.StringFrom("administrator"),
		Password:          string(hashed),
		Name:              "Admin",
		Surname:           "",
		Confirmed:         true,
	}

	err = userDao.Insert(ctx, newUser)

	if err != nil {
		return err
	}

	return nil
}

func downAdminUser(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
