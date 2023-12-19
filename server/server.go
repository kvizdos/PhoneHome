package server

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kvizdos/callhome/calls"
	"github.com/kvizdos/callhome/views"
	"github.com/rivo/tview"
)

func StartHTTP(app *tview.Application) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[1:]
		call, err := calls.GetCall(id)

		if err != nil {
			// handle
			return
		}

		call.Unanswered = true
		call.LastPhone = time.Now()

		if call.Requests == nil {
			call.Requests = make([]calls.Request, 0)
		}

		data := ""

		if r.Method == http.MethodGet && r.URL.Query().Has("data") {
			rawData := r.URL.Query().Get("data")
			// Allow for optional Base64 Encoding
			out, err := base64.StdEncoding.DecodeString(rawData)

			if err == nil {
				data = string(out)
			} else {
				data = rawData
			}
		}

		call.Requests = append(call.Requests, calls.Request{
			UserAgent:  r.Header.Get("user-agent"),
			Data:       data,
			Time:       time.Now(),
			HTTPMethod: r.Method,
		})

		calls.UpdateByID(call.ID, call)

		sendNotification(call.Name)

		app.QueueUpdateDraw(func() {
			views.ListView.SetItemText(call.Index, call.Name, fmt.Sprintf("Phoned Home @ %s", time.Now().Format("03:04 PM 02/01")))
		})

		fmt.Fprintf(w, "you've phoned home.")
	})

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
