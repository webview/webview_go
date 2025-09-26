package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("UA Test")
	w.SetSize(800, 600, webview.HintNone)

	w.SetUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 KemalTest/1.0")
	
	w.Navigate("https://www.whatismybrowser.com/detect/what-is-my-user-agent/")

	w.Run()
}
