VERSION := v1.48.0

init:
	-rm ./example.so
	-rm ./go.*
	go mod init
	-rm -r ${GOPATH}/pkg/mod/cache/download/github.com/golangci/golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@${VERSION}
	go mod tidy
	go build -buildmode=plugin plugin/example.go

update-golangci-lint:
	rm ./example.so
	rm ./go.*
	go mod init
	rm -r ${GOPATH}/pkg/mod/cache/download/github.com/golangci/golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@${VERSION}
	go mod tidy
	go build -buildmode=plugin plugin/example.go
