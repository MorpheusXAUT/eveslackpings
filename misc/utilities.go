package misc

import (
	"strings"
)

func ValidSlackToken(validTokens []string, token string) bool {
	for _, t := range validTokens {
		if strings.EqualFold(token, t) {
			return true
		}
	}

	return false
}
