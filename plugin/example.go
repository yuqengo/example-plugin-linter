// This must be package main
package main

import (
	"fmt"
	linters "github.com/dbraley/example-linter"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func main() {
	fmt.Printf("Plugin library for %s", AnalyzerPlugin.GetLinterName())
	fmt.Println("To compile this into a plugin, run:")
	fmt.Println("	 go build -buildmode=plugin -o /desired/path/to/example.so plugin/example.go")
	fmt.Println("See README.md for more info")
}

// This must be implemented
func (*analyzerPlugin) GetLinterName() string {
	return "example"
}

// This must be implemented
func (*analyzerPlugin) GetLinterDesc() string {
	return "A package of custom linters that can be used in golangci-lint"
}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.TodoAnalyzer,
	}
}

// This must be defined
var AnalyzerPlugin analyzerPlugin
