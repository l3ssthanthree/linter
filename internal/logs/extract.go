package logs

import "go/ast"

func ExtractMessage(call *ast.CallExpr) (ast.Expr, bool) {
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

	return call.Args[0], true
}
