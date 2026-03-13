package logs

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func ExtractMessage(pass *analysis.Pass, call *ast.CallExpr) (ast.Expr, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, false
	}

	switch sel.Sel.Name {
	case "Info", "Error", "Warn", "Debug":
	default:
		return nil, false
	}

	if len(call.Args) == 0 {
		return nil, false
	}

	if pkgIdent, ok := sel.X.(*ast.Ident); ok && pkgIdent.Name == "slog" {
		return call.Args[0], true
	}

	if isZapLogger(pass, sel.X) {
		return call.Args[0], true
	}

	return nil, false
}

func isZapLogger(pass *analysis.Pass, expr ast.Expr) bool {
	tv, ok := pass.TypesInfo.Types[expr]
	if !ok || tv.Type == nil {
		return false
	}

	return isZapLoggerType(tv.Type)
}

func isZapLoggerType(t types.Type) bool {
	switch tt := t.(type) {
	case *types.Pointer:
		return isZapLoggerNamed(tt.Elem())
	case *types.Named:
		return isZapLoggerNamed(tt)
	default:
		return false
	}
}

func isZapLoggerNamed(t types.Type) bool {
	named, ok := t.(*types.Named)
	if !ok {
		return false
	}

	obj := named.Obj()
	if obj == nil || obj.Pkg() == nil {
		return false
	}

	return obj.Name() == "Logger" && obj.Pkg().Path() == "go.uber.org/zap"
}
