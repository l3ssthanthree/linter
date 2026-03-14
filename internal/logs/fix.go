package logs

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
)

func LowercaseSuggestedText(expr ast.Expr) ([]byte, bool) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return nil, false
	}

	s, err := strconv.Unquote(lit.Value)
	if err != nil || s == "" {
		return nil, false
	}

	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToLower(r)
			break
		}
	}

	newQuoted := strconv.Quote(string(runes))
	return []byte(newQuoted), true
}
