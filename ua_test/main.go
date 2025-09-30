package main

import (
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetUserAgent("MySuperCustomUserAgent/1.0 (macOS)")
	w.Navigate("https://www.whatismybrowser.com/detect/what-is-my-user-agent/")
	w.Run()
}
