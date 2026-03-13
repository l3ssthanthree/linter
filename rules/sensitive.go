package rules

import (
	"regexp"
)

var sensitiveKeyWordsPatterns = []*regexp.Regexp{

	regexp.MustCompile(`(?i)\bpassword\b`), // password
	regexp.MustCompile(`(?i)\bpasswd\b`),   //passwd
	regexp.MustCompile(`(?i)\bpwd\b`),      // pwd

	regexp.MustCompile(`(?i)\token\b`),        // token
	regexp.MustCompile(`(?i)\bapi[_-]?key\b`), // api_key, api-key, apikey

	regexp.MustCompile(`(?i)\secret\b`), // secret

	regexp.MustCompile(`(?i)\bcredential(s)?\b`), // credential, credentials
}

// password: xxx
// password = xxx
// token: abc
// api_key=123
var sensitiveAssignmentPatterns = []*regexp.Regexp{

	regexp.MustCompile(`(?i)\bpassword\b\s*[:=]`),
	regexp.MustCompile(`(?i)\bpasswd\b\s*[:=]`),
	regexp.MustCompile(`(?i)\bpwd\b\s*[:=]`),

	regexp.MustCompile(`(?i)\btoken\b\s*[:=]`),

	regexp.MustCompile(`(?i)\bapi[_-]?key\b\s*[:=]`),

	regexp.MustCompile(`(?i)\bsecret\b\s*[:=]`),

	regexp.MustCompile(`(?i)\bcredential(s)?\b\s*[:=]`),
}

func HasSensitiveKeyword(s string) bool {
	for _, pattern := range sensitiveAssignmentPatterns {
		if pattern.MatchString(s) {
			return true
		}

		for _, pattern := range sensitiveKeyWordsPatterns {
			if pattern.MatchString(s) {
				return true
			}
		}
	}

	return false
}
