# webview bindings for Go

[![GoDoc](https://godoc.org/github.com/webview/webview_go?status.svg)](https://godoc.org/github.com/webview/webview_go)
[![Go Report Card](https://goreportcard.com/badge/github.com/webview/webview_go)](https://goreportcard.com/report/github.com/webview/webview_go)

### Getting Started

See [Go package documentation][go-docs] for the Go API documentation, or simply read the source code.

Start with creating a new directory structure for your project:

```sh
mkdir my-project && cd my-project
mkdir build
```

Create a new Go module:

```sh
go mod init example.com/m
```

Save the basic Go example into your project directory:

```sh
curl -sSLo basic.go "https://raw.githubusercontent.com/webview/webview_go/master/examples/basic/main.go"
```

Install dependencies:

```sh
go get github.com/webview/webview_go
```

Build and run the example:

```sh
# Linux, macOS
go build -o build/basic basic.go && ./build/basic
# Windows
go build -ldflags="-H windowsgui" -o build/basic.exe basic.go && "build/basic.exe"
```

### Notes

Calling `Eval()` or `Dispatch()` before `Run()` does not work because the webview instance has only been configured and not yet started.

[go-docs]:           https://pkg.go.dev/github.com/webview/webview
