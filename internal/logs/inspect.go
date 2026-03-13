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

func CollectStringLiterals(expr ast.Expr) []string {
	var result []string

	var walk func(ast.Expr)

	walk = func(e ast.Expr) {
		switch v := e.(type) {

		case *ast.BasicLit:
			if s, ok := StringLiteral(v); ok {
				result = append(result, s)
			}

		case *ast.BinaryExpr:
			walk(v.X)
			walk(v.Y)

		case *ast.CallExpr:
			for _, arg := range v.Args {
				walk(arg)
			}
		}
	}

	walk(expr)

	return result
}
