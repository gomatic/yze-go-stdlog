package stdlog

import "strconv"

// importPathLiteral is the raw text of an import spec's Path.Value: a Go string
// literal — interpreted ("log") or raw (`log`) — whose decoded content is the
// import path it denotes.
type importPathLiteral string

// importsLog reports whether the literal denotes the standard log package. Go
// permits an import path to be written as either an interpreted ("log") or a raw
// (`log`) string literal, and both bind the same package; strconv.Unquote decodes
// either form to the bare path, so a raw-literal import is matched as reliably as
// an interpreted one. A value that is not a valid Go string literal cannot denote
// any import and is rejected.
func importsLog(literal importPathLiteral) bool {
	path, err := strconv.Unquote(string(literal))
	return err == nil && path == stdLogPath
}
