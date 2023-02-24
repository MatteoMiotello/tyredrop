package utils

import (
	"github.com/friendsofgo/errors"
	"strconv"
	"strings"
)

func IsEmail(email string) bool {
	parts := strings.Split(email, "@")

	if len(parts) != 2 {
		return false
	}

	domain := parts[1]

	domainParts := strings.Split(domain, ".")

	if len(domainParts) < 2 {
		return false
	}

	return true
}

func FromStringToInt(stringPrice string, numberOfDecimals int) (int, error) {
	parts := strings.Split(stringPrice, ".")

	if len(parts) > 2 {
		return 0, errors.New("not valid number")
	}

	firstPart := parts[0]
	secondPart := parts[1]

	if len(secondPart) < numberOfDecimals {
		toPad := numberOfDecimals - len(secondPart)

		for i := 0; i < toPad; i++ {
			secondPart += "0"
		}
	}

	stringInt := firstPart + secondPart

	return strconv.Atoi(stringInt)
}
