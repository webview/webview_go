# webview_go

[![GoDoc](https://godoc.org/github.com/webview/webview_go?status.svg)](https://godoc.org/github.com/webview/webview_go)
[![Go Report Card](https://goreportcard.com/badge/github.com/webview/webview_go)](https://goreportcard.com/report/github.com/webview/webview_go)

Go language binding for the [webview library][webview].

> [!NOTE]
> Versions <= 0.1.1 are available in the [old repository][webview].

### Getting Started

See [Go package documentation][go-docs] for the Go API documentation, or simply read the source code.

Start with creating a new directory structure for your project.

```sh
mkdir my-project && cd my-project
```

Create a new Go module.

```sh
go mod init example.com/app
```

Save one of the example programs into your project directory.

```sh
curl -sSLo main.go "https://raw.githubusercontent.com/webview/webview_go/master/examples/basic/main.go"
```

Install dependencies.

```sh
go get github.com/webview/webview_go
```

Build the example. On Windows, add `-ldflags="-H windowsgui"` to the command line.

```sh
go build
```

### Notes

Calling `Eval()` or `Dispatch()` before `Run()` does not work because the webview instance has only been configured and not yet started.

[go-docs]: https://pkg.go.dev/github.com/webview/webview_go
[webview]: https://github.com/webview/webview


### WebView UserAgent Example

This project demonstrates how to set a custom User-Agent string in webview_go.

It builds a simple Go application that shows the current User-Agent in the UI and verifies that it matches the value defined in the source code.

Requirements
Go 1.18+
C compiler (clang/gcc)
Linux:
libgtk-3-dev
libwebkit2gtk-4.0-dev
macOS: no additional dependencies

### Build 

chmod +x build.sh
./build.sh

This will:
	1.	Install necessary dependencies (on Linux).
	2.	Compile main.go into an executable webview_app.

### Run 

./webview_app

Expected behavior

When the application runs:
	•	A window will open with:
	•	A UserAgent section showing the custom value (e.g. MyCustomUserAgent/1.0).
	•	A button and counter to demonstrate bindings.

If you see your custom User-Agent displayed in the window, the implementation is correct

Troubleshooting

Linux:
1. Error: webkit2/webkit2.h: No such file or directory
-> Install missing GTK/WebKit dependencies:
sudo apt-get update
sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.0-dev

2. Error: pkg-config not found
-> Install pkg-config:
sudo apt-get install -y pkg-config

3. Blank window / crash on startup
-> Ensure you are running inside a desktop environment with GTK/WebKit support.
-> On minimal VMs, install ubuntu-desktop or run inside an X11 session.

MacOS:
1. Error: ld: framework not found WebKit
-> Ensure you use clang (not gcc) when compiling.
-> Go automatically links against Cocoa/WebKit frameworks, so usually no manual fix is required.

2. Window opens but is empty
-> Sometimes macOS blocks WebKit initialization in sandboxed terminals. Try:
open ./webview_test

3. Custom UserAgent not applied
-> Make sure you call w.SetUserAgent("MyCustomUserAgent/1.0") before w.Navigate(...) or w.SetHtml(...)

