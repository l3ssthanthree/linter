package analyzer

import (
	"go/ast"
	"mylinter/internal/logs"
	"mylinter/rules"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			msgExpr, ok := logs.ExtractMessage(call)
			if !ok {
				return true
			}

			msg, ok := logs.StringLiteral(msgExpr)
			if !ok {
				return true
			}

			if !rules.StartsWithLowercase(msg) {
				pass.Reportf(msgExpr.Pos(), "log message must start with a lowercase letter")
			}

			if !rules.IsEnglishOnly(msg) {
				pass.Reportf(msgExpr.Pos(), "log message must contain only English text")
			}

			return true
		})
	}
	return nil, nil
}
