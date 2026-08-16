// Package stdlog provides a go/analysis analyzer enforcing the gomatic Go logging
// standard: the standard library "log" package is forbidden in favor of
// structured logging with "log/slog".
package stdlog

import (
	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
)

const (
	message = "the standard log package is forbidden; use log/slog for structured logging"
	// stdLogPath is the import path of the standard log package, as it appears
	// once an ast.ImportSpec's quoted Path.Value has been decoded.
	stdLogPath = "log"
)

// Analyzer reports imports of the standard "log" package.
var Analyzer = &analysis.Analyzer{
	Name: "stdlog",
	Doc:  "reports imports of the standard log package, which the gomatic standard replaces with log/slog",
	Run:  run,
}

// Registration declares this analyzer to the yze framework. The category is
// "logging", shared with yze/slogkv, so narrowing a run to the logging standard
// yields the whole standard rather than half of it; "data" was this rule's tag
// until then and named the wrong subject — it is jsontag's, about serialization.
var Registration = goyze.Registration{
	Name:       "stdlog",
	Categories: []goyze.Category{"logging"},
	URL:        "https://docs.gomatic.dev/yze/stdlog",
	Analyzer:   Analyzer,
}

// run reports each import of the standard log package.
func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			if importsLog(importPathLiteral(imp.Path.Value)) {
				pass.Reportf(imp.Pos(), message)
			}
		}
	}
	return nil, nil
}
