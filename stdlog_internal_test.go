package stdlog

import "testing"

// TestImportsLog pins the contract of importsLog directly, with Go string
// literals, so the raw-import case ("`log`") is covered without a raw-string
// import in testdata (which gofumpt would rewrite to "log", erasing the test).
func TestImportsLog(t *testing.T) {
	cases := []struct {
		name    string
		literal importPathLiteral
		want    bool
	}{
		{name: "interpreted log", literal: `"log"`, want: true},
		{name: "raw log", literal: "`log`", want: true},
		{name: "interpreted log/slog", literal: `"log/slog"`, want: false},
		{name: "interpreted fmt", literal: `"fmt"`, want: false},
		{name: "unquoted log", literal: "log", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := importsLog(tc.literal); got != tc.want {
				t.Errorf("importsLog(%q) = %t, want %t", tc.literal, got, tc.want)
			}
		})
	}
}
