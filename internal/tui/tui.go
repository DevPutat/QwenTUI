package tui

import (
	"github.com/DevPutat/QwenTUI/internal/request"
	"github.com/DevPutat/QwenTUI/internal/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewApp(conf *types.Conf) {
	tApp := tview.NewApplication()
	tInputChatField := tview.NewInputField().
		SetLabel("Введите запрос: ").
		SetFieldWidth(50)

	tOutputField := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			tApp.Draw()
		})
	app := &types.App{
		App:       tApp,
		ChatField: tOutputField,
	}

	sendButton := tview.NewButton("Отправить").
		SetSelectedFunc(func() {
			query := tInputChatField.GetText()
			if query == "" {
				return
			}
			tInputChatField.SetText("")
			tOutputField.Clear()
			go request.SendStream(query, conf, app)
		})

	grid := tview.NewGrid().
		SetRows(0, 3).
		SetColumns(0).
		AddItem(tInputChatField, 0, 0, 1, 1, 0, 0, true).
		AddItem(sendButton, 1, 0, 1, 1, 0, 0, false).
		AddItem(tOutputField, 2, 0, 1, 1, 0, 0, false)

	focusManger := tview.NewFlex().AddItem(grid, 0, 1, true)
	curFocus := 0
	elements := []tview.Primitive{tInputChatField, sendButton}

	tApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			curFocus = (curFocus + 1) & len(elements)
			app.App.SetFocus(elements[curFocus])
			return nil
		case tcell.KeyBacktab:
			curFocus = (curFocus - 1 + len(elements)) & len(elements)
			app.App.SetFocus(elements[curFocus])
			return nil
		}
		return event
	})

	if err := tApp.SetRoot(focusManger, true).SetFocus(tInputChatField).Run(); err != nil {
		panic(err)
	}
}
