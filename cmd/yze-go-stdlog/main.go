// Command yze-go-stdlog runs the stdlog analyzer as a standalone go/analysis
// checker (text and -json output, and as a `go vet -vettool`).
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	stdlog "github.com/gomatic/yze-go-stdlog"
)

// run is the analysis entry point, indirected so the binary's wiring is testable
// without invoking the real driver (which loads packages and exits the process).
var run = singlechecker.Main

func main() { run(stdlog.Analyzer) }
