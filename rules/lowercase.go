package rules

import "unicode"

func StartsWithLowercase(s string) bool {
	if s == "" {
		return true
	}

	for _, r := range s {
		if unicode.IsLetter(r) {
			return unicode.IsLower(r)
		}
	}

	return true
}
