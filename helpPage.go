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

	boxBody := tview.NewTable()
	boxBody.SetBorder(true)
	for i := range ipanelMenu {
		boxBody.SetCellSimple(i*2, 0, ipanelMenu[i])
		boxBody.GetCell(i*2, 0).SetAlign(tview.AlignLeft).SetExpansion(7)
		boxBody.SetCellSimple(i*2+1, 0, "review option")
		boxBody.GetCell(i*2+1, 0).SetAlign(tview.AlignCenter).SetExpansion(3)
	}

	helpPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxHelp, 3, 1, false).
		AddItem(boxBody, 0, 3, false)
	return helpPage
}
