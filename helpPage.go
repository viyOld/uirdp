package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func getHelpPage() *tview.Flex {
	// helpPage
	boxHelp := tview.NewTextView()
	boxHelp.SetBorder(true)
	boxHelp.SetText("Help Page")
	boxHelp.SetTextAlign(tview.AlignCenter)
	boxHelp.SetBackgroundColor(tcell.Color(987))
	helpPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxHelp, 5, 3, false)
	return helpPage
}
