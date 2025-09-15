package main

import webview "github.com/webview/webview_go"

const html = `<h2>UserAgent:</h2>
<div id="ua"></div>
<button id="increment">Tap me</button>
<div>You tapped <span id="count">0</span> time(s).</div>
<script>
  const [incrementElement, countElement] =
    document.querySelectorAll("#increment, #count");

  document.addEventListener("DOMContentLoaded", () => {
    document.getElementById("ua").textContent = navigator.userAgent;

    incrementElement.addEventListener("click", () => {
      window.increment().then(result => {
        countElement.textContent = result.count;
      });
    });
  });
</script>`

type IncrementResult struct {
	Count uint `json:"count"`
}

func main() {
	var count uint = 0
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Bind and UserAgent Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetUserAgent("MyCustomUserAgent/1.0")

	// A binding that increments a value and immediately returns the new value.
	w.Bind("increment", func() IncrementResult {
		count++
		return IncrementResult{Count: count}
	})

	w.SetHtml(html)
	w.Run()
}
