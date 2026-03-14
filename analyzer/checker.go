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

			msgExpr, ok := logs.ExtractMessage(pass, call)
			if !ok {
				return true
			}

			if msg, ok := logs.StringLiteral(msgExpr); ok {
				if !rules.StartsWithLowercase(msg) {
					if newText, ok := logs.LowercaseSuggestedText(msgExpr); ok {
						pass.Report(analysis.Diagnostic{
							Pos:     msgExpr.Pos(),
							End:     msgExpr.End(),
							Message: "log message must start with a lowercase letter",
							SuggestedFixes: []analysis.SuggestedFix{
								{
									Message: "convert first letter to lowercase",
									TextEdits: []analysis.TextEdit{
										{
											Pos:     msgExpr.Pos(),
											End:     msgExpr.End(),
											NewText: newText,
										},
									},
								},
							},
						})
					} else {
						pass.Reportf(msgExpr.Pos(), "log message must start with a lowercase letter")
					}
				}

				if !rules.IsEnglishOnly(msg) {
					pass.Reportf(msgExpr.Pos(), "log message must contain only English text")
				}

				if rules.HasForbiddenSymbols(msg) {
					pass.Reportf(msgExpr.Pos(), "log message must not contain special symbols or emoji")
				}

				if rules.HasSensitiveKeyword(msg) {
					pass.Reportf(msgExpr.Pos(), "log message may contain sensitive data")
				}
			}

			for _, s := range logs.CollectStringLiterals(msgExpr) {
				if rules.HasSensitiveKeyword(s) {
					pass.Reportf(msgExpr.Pos(), "log message may contain sensitive data")
				}
			}

			return true
		})
	}
	return nil, nil
}
