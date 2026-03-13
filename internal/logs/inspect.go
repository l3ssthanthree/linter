package logs

import (
	"go/ast"
	"go/token"
	"strconv"
)

func StringLiteral(expr ast.Expr) (string, bool) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok {
		return "", false
	}

	if lit.Kind != token.STRING {
		return "", false
	}

	s, err := strconv.Unquote(lit.Value)
	if err != nil {
		return "", false
	}

	return s, true
}
