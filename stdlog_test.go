package stdlog_test

import (
	"testing"

	goyze "github.com/gomatic/go-yze"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"

	stdlog "github.com/gomatic/yze-go-stdlog"
)

func TestStandardLogImportIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), stdlog.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, stdlog.Registration.Validate())
	assert.Equal(t, "yze/stdlog", stdlog.Registration.RuleID())
	assert.Same(t, stdlog.Analyzer, stdlog.Registration.Analyzer)
}

// TestRegistrationCategoryIsLogging pins the category, which nothing else can
// fail on: goyze.Category is a bare string, Registration.Validate checks only
// Name and Analyzer, and a corpus Finding carries no category at all. "logging"
// is the tag this rule shares with yze/slogkv, so a consumer narrowing a run to
// the logging standard gets the whole standard rather than half of it.
func TestRegistrationCategoryIsLogging(t *testing.T) {
	assert.Equal(t, []goyze.Category{"logging"}, stdlog.Registration.Categories)
}
