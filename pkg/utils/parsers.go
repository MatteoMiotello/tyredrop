package utils

import "strings"

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
