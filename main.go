package main

import (
	"github.com/kvizdos/callhome/server"
	"github.com/kvizdos/callhome/views"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	go server.StartHTTP(app)

	primaryAppLayout := views.GeneratePrimaryView(app)

	// Set root and start the application.
	if err := app.SetRoot(primaryAppLayout, true).Run(); err != nil {
		panic(err)
	}
}
