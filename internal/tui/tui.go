package tui

import (
	"github.com/DevPutat/QwenTUI/internal/request"
	"github.com/DevPutat/QwenTUI/internal/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewApp(conf *types.Conf) {
	tApp := tview.NewApplication().EnableMouse(true)
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

	// sendButton := tview.NewButton("Отправить").
	// 	SetSelectedFunc(func() {
	// 		query := tInputChatField.GetText()
	// 		if query == "" {
	// 			return
	// 		}
	// 		tInputChatField.SetText("")
	// 		tOutputField.Clear()
	// 		go request.SendStream(query, conf, app)
	// 	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tOutputField, 0, 5, false)
	flex.AddItem(tInputChatField, 0, 1, true)

	tApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		// case tcell.KeyTab:
		// 	curFocus = (curFocus + 1) & len(elements)
		// 	tApp.SetFocus(elements[curFocus])
		// 	return nil
		case tcell.KeyEnter:
			query := tInputChatField.GetText()
			if query == "" {
				return nil
			}
			tInputChatField.SetText("")
			tOutputField.Clear()
			go request.SendStream(query, conf, app)

			return nil
		}
		return event
	})

	if err := tApp.SetRoot(flex, true).SetFocus(tInputChatField).Run(); err != nil {
		panic(err)
	}
}
