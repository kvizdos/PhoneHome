package views

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kvizdos/callhome/calls"
	"github.com/rivo/tview"
)

func GenerateDetailsOverviewView(app *tview.Application, call calls.Call) tview.Primitive {
	list := tview.NewList()
	list.SetSecondaryTextColor(tcell.Color39)
	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		detailsView := GenerateDetailedRequestView(app, call, call.Requests[len(call.Requests)-index-1])
		app.SetRoot(detailsView, true)
	})

	for i := len(call.Requests) - 1; i >= 0; i-- {
		req := call.Requests[i]
		list.AddItem(fmt.Sprintf("%s %s", req.Time.Format("03:04 PM 02/01"), req.HTTPMethod), fmt.Sprintf("Payload: %d bytes", len(req.Data)), 0, nil)
	}

	// Monitor for new Requests
	lastSeenNumRequests := len(call.Requests)
	go func(list *tview.List, getLatestCall func() calls.Call) {
		for {
			call := getLatestCall()
			if lastSeenNumRequests != len(call.Requests) {
				app.QueueUpdateDraw(func() {
					list.Clear()
				})
				for i := len(call.Requests) - 1; i >= 0; i-- {
					req := call.Requests[i]
					app.QueueUpdateDraw(func() {
						list.AddItem(fmt.Sprintf("%s %s", req.Time.Format("03:04 PM 02/01"), req.HTTPMethod), fmt.Sprintf("Payload: %d bytes", len(req.Data)), 0, nil)
					})
				}
				lastSeenNumRequests = len(call.Requests)
			}
			time.Sleep(1 * time.Second)
		}
	}(list, func() calls.Call {
		newCall, _ := calls.GetCall(call.ID)
		call = newCall
		return newCall
	})

	averageDelay := call.AverageRequestDelay()
	var formattedDelay string

	if averageDelay.Seconds() < 60 {
		formattedDelay = fmt.Sprintf("%.1f seconds", averageDelay.Seconds())
	} else {
		formattedDelay = fmt.Sprintf("%.1f minutes", averageDelay.Minutes())
	}
	callInfoView := tview.NewTextView().SetText(fmt.Sprintf("%s - Call Record\nURL ID: %s\nFrequency: %s", call.Name, call.ID, formattedDelay))

	grid := tview.NewGrid().
		SetRows(0).         // One row that takes up all vertical space
		SetColumns(-1, -2). // Two columns: -2 (2/3 of the width), -1 (1/3 of the width)
		SetBorders(true)    // Optional: adds borders for visualization

	grid.AddItem(callInfoView, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(list, 0, 1, 1, 1, 0, 0, false)

	// InputField for command input.
	labelColorFocused := tcell.Color39
	labelColorBlurred := tcell.Color57
	inputField := tview.NewInputField()
	inputField.SetLabel("Command : ").
		SetLabelColor(labelColorFocused).
		SetFieldWidth(0).
		SetDoneFunc(func(key tcell.Key) {
			// Handle command input.
			command := inputField.GetText()
			if command == "q" {
				app.SetRoot(GeneratePrimaryView(app), true)
				app.SetFocus(ListView)
				return
			}
			// Clear input field and focus on it again.
			inputField.SetText("")
			app.SetFocus(inputField)
		})

	inputField.SetFocusFunc(func() {
		inputField.SetLabelColor(labelColorFocused)
	})

	inputField.SetBlurFunc(func() {
		inputField.SetLabelColor(labelColorBlurred)
	})

	// Layout for our app.
	primaryAppLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(grid, 0, 1, false).
		AddItem(inputField, 1, 1, true)

	// Handle Key Presses & Controls
	primaryAppLayout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			app.SetFocus(inputField)
			return nil
		}
		if event.Key() == tcell.KeyEscape {
			app.SetFocus(list)
			return nil
		}
		return event
	})

	app.SetFocus(list)

	return primaryAppLayout
}
