package analyzer

import (
	"go/ast"
	"mylinter/internal/logs"

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

			pass.Reportf(msgExpr.Pos(), "found log message: %q", msg)
			_ = msg

			return true
		})
	}
	return nil, nil
}
