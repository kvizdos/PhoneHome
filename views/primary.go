package views

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kvizdos/callhome/calls"
	"github.com/rivo/tview"
)

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var ListView = tview.NewList()

func GeneratePrimaryView(app *tview.Application) tview.Primitive {
	ListView.SetBorder(true).SetTitle(" Contact Manifest ") // Optional: Adjust or remove
	ListView.SetBorderPadding(0, 0, 0, 0)
	ListView.SetSecondaryTextColor(tcell.Color39)
	ListView.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		call, _ := calls.GetCallByIndex(index)
		detailedView := GenerateDetailsOverviewView(app, call)
		app.SetRoot(detailedView, true)
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
			if strings.HasPrefix(command, "new ") {
				newItem := strings.TrimPrefix(command, "new ")
				newItem = strings.TrimSpace(newItem)
				id := generateRandomString(8)
				calls.Calls = append(calls.Calls, calls.Call{
					Name:  newItem,
					Index: len(calls.Calls),
					ID:    id,
				})
				ListView.AddItem(newItem, fmt.Sprintf("%s - Waiting...", id), 0, nil)
			}
			if strings.HasPrefix(command, "update ") {
				newItem := strings.TrimPrefix(command, "update ")
				ListView.SetItemText(0, newItem, "")
			}
			// Clear input field and focus on it again.
			inputField.SetText("")
			app.SetFocus(ListView)
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
		AddItem(ListView, 0, 1, false).
		AddItem(inputField, 1, 1, true)

	// Handle Key Presses & Controls
	primaryAppLayout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			app.SetFocus(inputField)
			return nil
		}
		if event.Key() == tcell.KeyEscape {
			app.SetFocus(ListView)
			return nil
		}

		if event.Rune() == 'r' && ListView.HasFocus() {
			currentIndex := ListView.GetCurrentItem()
			call, _ := calls.GetCallByIndex(currentIndex)
			ListView.SetItemText(currentIndex, call.Name, fmt.Sprintf("%s - Waiting...", call.ID))
		}
		return event
	})

	return primaryAppLayout
}
