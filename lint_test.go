package linters

// Tests for linters.

import (
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"runtime"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

type linterSuite struct {
	suite.Suite
}

func (suite *linterSuite) TestContextLinter() {
	analysistest.Run(
		suite.T(), TestdataDir(),
		TodoAnalyzer, "testlintdata/todo")
}

func TestLinterSuite(t *testing.T) {
	suite.Run(t, new(linterSuite))
}

func TestdataDir() string {
	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get current test filename")
	}
	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
