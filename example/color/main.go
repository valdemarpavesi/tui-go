package main

import (
	"github.com/marcusolsson/tui-go"
)

func main() {
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems("First row", "Second row", "Third row", "Fourth row", "Fifth row", "Sixth row")
	l.SetSelected(0)

	warning := tui.NewLabel("WARNING: This is a warning")
	warning.SetStyleName("warning")

	fatal := tui.NewLabel("FATAL: Cats and dogs are now living together.")
	fatal.SetStyleName("fatal")

	root := tui.NewVBox(l, warning, fatal)

	t := tui.NewTheme()
	t.SetStyle("list.item", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorMagenta})
	t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorRed, Fg: tui.ColorWhite})

	// The style name is appended to the widget name to support coloring of
	// individual labels.
	t.SetStyle("label.fatal", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})
	t.SetStyle("label.warning", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorYellow})

	ui := tui.New(root)
	ui.SetTheme(t)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
