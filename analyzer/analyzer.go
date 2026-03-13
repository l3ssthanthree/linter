package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "logcheck",
	Doc:  "check log messages in slog and zap",
	Run:  run,
}
