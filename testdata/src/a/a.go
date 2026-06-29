package a

import (
	"log" // want `standard log package`
	"log/slog"
)

var _ = log.Println
var _ = slog.Info
