package main

import webview "github.com/webview/webview_go"

const html = `<html style="background: #1B2845; color: #eee;">
<button onclick="increment();">Tap me</button>
<div>You tapped <span id="count">0</span> time(s).</div>
<script>
	const counter = document.getElementById("count")
	async function increment() {
		const result = await window.Increment()
		counter.textContent = result.count;
	}
</script>
</html>`

type IncrementResult struct {
	Count uint `json:"count"`
}

func main() {
	var count uint = 0
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(480, 320, webview.HintNone)

	// A binding that increments a value and immediately returns the new value.
	w.Bind("Increment", func() IncrementResult {
		count++
		return IncrementResult{Count: count}
	})

	w.SetHtml(html)
	w.Run()
}
