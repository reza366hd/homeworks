package main

import "github.com/marcusolsson/tui-go"

func main() {
	box := tui.NewVBox(
		tui.NewLabel("Hello world"),
	)

	ui, _ := tui.New(box)

	ui.SetKeybinding("Tabs", func() { ui.Quit() })
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	ui.Run()
}
