package app

import "github.com/coding-for-fun-org/chat-gpt-tui/pkg/tui"

func Start() {
	app := tui.NewTUI()

	app.Run()
}
