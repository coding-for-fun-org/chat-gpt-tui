package tui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	application *tview.Application
	// chatGroupView *tview.Flex
	chatLogView   *tview.TextView
	chatInputView *tview.TextArea
	mainView      *tview.Flex
}

func NewTUI() *TUI {
	tui := &TUI{
		application:   nil,
		chatLogView:   nil,
		chatInputView: nil,
	}

	tui.application = tui.initApp()
	tui.chatLogView = tui.initChatLogView()
	tui.chatInputView = tui.initChatInputView()
	tui.mainView = tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Chat Groups"), 40, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tui.chatLogView, 0, 1, false).
			AddItem(tui.chatInputView, 7, 1, false), 0, 1, false)

	return tui
}

func (tui *TUI) Run() {
	if err := tui.application.SetRoot(tui.mainView, true).SetFocus(tui.mainView).Run(); err != nil {
		panic(err)
	}
}

func (tui *TUI) initApp() *tview.Application {
	application := tview.NewApplication()

	application.EnableMouse(true)

	return application
}

func (tui *TUI) initChatLogView() *tview.TextView {
	textView := tview.NewTextView()

	textView.SetTitle("Chats")
	textView.SetBorder(true)
	textView.SetDynamicColors(true)
	textView.SetRegions(true)
	textView.SetChangedFunc(func() {
		tui.application.Draw()
	})

	return textView
}

func (tui *TUI) initChatInputView() *tview.TextArea {
	textArea := tview.NewTextArea()

	textArea.SetPlaceholder("Enter text here...")
	textArea.SetTitle("Text Area")
	textArea.SetBorder(true)
	textArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			log.Println("Enter pressed")

			return nil
		}

		return event
	})
	textArea.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
		return action, event
	})

	return textArea
}
