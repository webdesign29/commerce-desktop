package main

import (
	"github.com/webview/webview"
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Application E-repas 1.0")
	w.Bind("printPOS", func(obj string) string {
		return "I am now printing stuff"
	})
	w.SetSize(1200, 800, webview.HintNone)
	w.Navigate("https://gestion.e-repas.com")
	w.Run()
}