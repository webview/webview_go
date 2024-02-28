package main

import webview "github.com/webview/webview_go"

const html = `<button id="increment">Tap me</button>
<div>You tapped <span id="count">0</span> time(s).</div>
<script>
  const incrementBtn = document.getElementById("increment");
  const counter = document.getElementById("count");
  incrementBtn.addEventListener("click", async () => {
    const result = await window.increment();
    counter.textContent = result.count;
  });
</script>`

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
	w.Bind("increment", func() IncrementResult {
		count++
		return IncrementResult{Count: count}
	})

	w.SetHtml(html)
	w.Run()
}
