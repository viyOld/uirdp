package main

import (
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

	headerBoxHelp := tview.NewTextView()
	headerBoxHelp.SetBorder(true)
	headerBoxHelp.SetDynamicColors(true)
	headerBoxHelp.SetText("[#8080ff] Help Page [::]")
	headerBoxHelp.SetTextAlign(tview.AlignCenter)
	//headerBoxHelp.SetTextColor(tcell.ColorGreen)

	helpBodyTable := tview.NewTable()
	helpBodyTable.SetBorder(true)
	helpBodyTable.SetBorderPadding(2, 2, 3, 3)
	for i := range ipanelMenu {
		helpBodyTable.SetCellSimple(i*3, 0, ipanelMenu[i])
		helpBodyTable.GetCell(i*2, 0).SetAlign(tview.AlignLeft).SetExpansion(7)
		helpBodyTable.SetCellSimple(i*3+1, 0, helpLine[i])
		helpBodyTable.GetCell(i*2+1, 0).SetAlign(tview.AlignLeft).SetExpansion(3)
	}
	//setFocus := app.SetFocus(helpBodyTable)
	helpPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerBoxHelp, 3, 1, false).
		AddItem(helpBodyTable, 0, 3, false)

	//helpPage.InputHandler()(event, app.SetFocus(helpBodyTable))

	return helpPage
}
