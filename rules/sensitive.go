package rules

import (
	"regexp"
)

var sensitiveKeywordPatterns = []*regexp.Regexp{

	regexp.MustCompile(`(?i)\bpassword\b`), // password
	regexp.MustCompile(`(?i)\bpasswd\b`),   //passwd
	regexp.MustCompile(`(?i)\bpwd\b`),      // pwd

	regexp.MustCompile(`(?i)\btoken\b`),              // token
	regexp.MustCompile(`(?i)\bapi(?:[_\-\s]?key)\b`), // api_key, api-key, apikey, api key

	regexp.MustCompile(`(?i)\bsecret\b`), // secret

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

	regexp.MustCompile(`(?i)\bapi(?:[_\-\s]?key)\b\s*[:=]`),

	regexp.MustCompile(`(?i)\bsecret\b\s*[:=]`),

	regexp.MustCompile(`(?i)\bcredential(s)?\b\s*[:=]`),
}

var customSensitivePatterns []*regexp.Regexp

func SetCustomSensitivePatterns(patterns []string) error {
	compiled := make([]*regexp.Regexp, 0, len(patterns))

	for _, pattern := range patterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}
		compiled = append(compiled, re)
	}

	customSensitivePatterns = compiled
	return nil
}

func HasSensitiveKeyword(s string) bool {
	for _, pattern := range sensitiveAssignmentPatterns {
		if pattern.MatchString(s) {
			return true
		}
	}

	for _, pattern := range sensitiveKeywordPatterns {
		if pattern.MatchString(s) {
			return true
		}
	}

	for _, pattern := range customSensitivePatterns {
		if pattern.MatchString(s) {
			return true
		}
	}

	return false
}
