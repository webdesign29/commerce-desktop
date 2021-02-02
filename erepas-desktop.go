package commerce-desktop

import (
	"github.com/webview/webview"
	"log"
	"runtime"

	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
)

func main() {
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/mouuff/go-rocket-update-example",
			ZipName:       "binaries_" + runtime.GOOS + ".zip",
		},
		BinaryName: "go-rocket-update-example",
		Version:    "v1.0.0", // You can change this value to trigger an update
	}
	log.Println(u.Version)
	err := u.Update()
	if err != nil {
		log.Fatal(err)
	}
	
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Application E-repas")
	w.Bind("printPOS", func(obj string) string {
		return "I am now printing stuff"
	})
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://gestion.e-repas.com")
	w.Run()
}



func add(a float64, b float64) float64 {
	return a + b
}
