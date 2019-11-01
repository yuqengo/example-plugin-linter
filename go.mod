module github.com/dbraley/example-linter

// All versions here need to be the same as in golangci-lint/mod.go if present

go 1.12

require (
	github.com/stretchr/testify v1.4.0
	golang.org/x/tools v0.0.0-20191010075000-0337d82405ff
)

replace golang.org/x/tools => github.com/golangci/tools v0.0.0-20190915081525-6aa350649b1c
