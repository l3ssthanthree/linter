package rules

import "unicode"

func IsEnglishOnly(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Cyrillic) {
			return false
		}
	}

	return true
}
