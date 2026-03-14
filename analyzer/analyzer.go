package analyzer

import "golang.org/x/tools/go/analysis"

var configPath string

var Analyzer = &analysis.Analyzer{
	Name: "logcheck",
	Doc:  "check log messages in slog and zap",
	Run:  run,
}

func init() {
	Analyzer.Flags.StringVar(&configPath, "config", ".logcheck.yml", "path to logcheck config file")
}
