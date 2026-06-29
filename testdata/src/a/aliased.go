package a

import (
	_ "log"      // want `standard log package`
	applog "log" // want `standard log package`
)

var _ = applog.Println
