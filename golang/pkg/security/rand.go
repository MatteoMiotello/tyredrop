package security

import (
	"crypto/rand"
	"github.com/spf13/viper"
)

func GenerateSecureKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)

	if err != nil {
		return "", err
	}

	chars := viper.GetString("APPLICATION_KEY")

	for i, b := range key {
		key[i] = chars[b%byte(len(chars))]
	}

	return string(key), nil
}
