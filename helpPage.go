package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var helpLine = []string{
	"The window where it is more developed explains the possibilities of the program and individual options",
	"A window showing the main system parameters related to the operation of the program",
	"Additional commands that are not part of the main menu",
	"Check the existence of a remote server and the ability to connect to it",
	"Initiating a session",
	"Create a new record",
	"Create a new record by copying the active one and editing it",
	"Edit an entry",
	"Deleting an entry",
	"Exit the current window. If this is the main window then exit the program.",
}

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
		boxBody.SetCellSimple(i*2+1, 0, helpLine[i])
		boxBody.GetCell(i*2+1, 0).SetAlign(tview.AlignCenter).SetExpansion(3)
	}

	helpPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxHelp, 3, 1, false).
		AddItem(boxBody, 0, 3, false)
	return helpPage
}
