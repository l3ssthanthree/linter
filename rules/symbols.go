package rules

import (
	"strings"
	"unicode"
)

func HasForbiddenSymbols(s string) bool {
	if strings.Contains(s, "...") {
		return true
	}

	for _, r := range s {
		switch r {
		case '!', '?':
			return true
		}

		if unicode.IsSymbol(r) {
			return true
		}
	}

	return false
}
