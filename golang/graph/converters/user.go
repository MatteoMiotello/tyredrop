package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func UserToGraphQL(user *models.User) *model.User {
	return &model.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: &user.Username.String,
	}
}
