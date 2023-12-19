package views

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kvizdos/callhome/calls"
	"github.com/rivo/tview"
)

func GenerateDetailedRequestView(app *tview.Application, call calls.Call, request calls.Request) tview.Primitive {
	headerText := fmt.Sprintf("[#00afff]Request Details for %s @ %s[-:-:-]\n", call.Name, request.Time.Format("03:04 PM 02/01"))

	combinedText := headerText

	combinedText += fmt.Sprintf("Method: [#ffff00]%s[-:-:-]\n", request.HTTPMethod)
	combinedText += fmt.Sprintf("User Agent: [#ffff00]%s[-:-:-]\n", request.UserAgent)

	dataValue := request.Data
	if len(dataValue) == 0 {
		dataValue = "No data sent."
	}
	combinedText += fmt.Sprintf("Data: [#ffff00]%s[-:-:-]", dataValue)

	textView := tview.NewTextView().
		SetText(combinedText).
		SetScrollable(true).
		SetDynamicColors(true) // Enable dynamic color rendering

	var currentLine int

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if currentLine > 0 {
				currentLine--
			}
		case tcell.KeyDown:
			currentLine++
		}

		// Update the TextView's content offset to scroll
		textView.ScrollTo(currentLine, 0)
		return event
	})
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
				app.SetRoot(GenerateDetailsOverviewView(app, call), true)
				return
			}
			if command == "dump data" {
				err := request.DumpDataToFile(call.Name, ".")
				if err != nil {
					// handle
				}
			}
			if command == "dump" {
				err := request.DumpToFile(call.Name, ".")
				if err != nil {
					// handle
				}
			}
			if command == "dump ugly" {
				err := request.DumpToUglyFile(call.Name, ".")
				if err != nil {
					// handle
				}
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
	primaryAppLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 0, 1, false).
		AddItem(inputField, 1, 1, true)

	// Handle Key Presses & Controls
	primaryAppLayout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			app.SetFocus(inputField)
			return nil
		}
		if event.Key() == tcell.KeyEscape {
			app.SetFocus(textView)
		}
		return event
	})

	app.SetFocus(inputField)

	return primaryAppLayout
}
