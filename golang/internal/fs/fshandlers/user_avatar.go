package fshandlers

import (
	"mime"
	"pillowww/titw/models"
	"strconv"
)

type UserAvatarFs struct {
	*fsHandler
}

func NewUserAvatar() *UserAvatarFs {
	return &UserAvatarFs{
		fsHandler: &fsHandler{
			BasePath:   rootPath + "/avatar",
			PublicPath: "/private/avatar",
		},
	}
}

func (u UserAvatarFs) StoreAvatar(user *models.User, ct string, stream []byte) (*string, error) {
	ext, err := mime.ExtensionsByType(ct)

	if err != nil {
		return nil, err
	}

	fileName := strconv.FormatInt(user.ID, 16) + ext[len(ext)-1]

	err = u.WriteFile(fileName, stream)

	if err != nil {
		return nil, err
	}

	return &fileName, nil
}
