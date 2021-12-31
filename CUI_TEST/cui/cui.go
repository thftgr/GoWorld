package cui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

var g *gocui.Gui

func Load() (gg *gocui.Gui, err error) {
	g, err = gocui.NewGui(gocui.OutputNormal)
	g.ASCII = true
	if err != nil {
		return
	}

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	g.SetManagerFunc(miniSizeHelp)

	if err := g.SetKeybinding("", gocui.KeyCtrlH, gocui.ModNone, toggleHelp); err != nil {
		return
	}

	gg = g
	return
}

func help(g *gocui.Gui) error {
	maxX, _ := g.Size()
	v, err := g.SetView("help", maxX-25, 0, maxX-1, 9)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "    KEYBINDINGS    ")
		fmt.Fprintln(v, "Space: New View")
		fmt.Fprintln(v, "Tab: Next View")
		fmt.Fprintln(v, "← ↑ → ↓: Move View")
		fmt.Fprintln(v, "Backspace: Delete View")
		fmt.Fprintln(v, "t: Set view on top")
		fmt.Fprintln(v, "b: Set view on bottom")
		fmt.Fprintln(v, "^C: Exit")
	}
	return nil
}
