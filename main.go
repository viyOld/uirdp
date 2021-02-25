package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app   *tview.Application // The tview application.
	pages *tview.Pages       // The application pages.

	ipanelMenu = []string{"h|Help", "i|Info", "m|Menu", "t|Test", "s|Start", "n|New", "c|Copy", "e|Edit", "d|Dellete", "q|Exit"}
)

func main() {
	app = tview.NewApplication()
	pages := tview.NewPages()

	// helpPage
	boxHelp := tview.NewTextView()
	boxHelp.SetBorder(true)
	boxHelp.SetText("Help Page")
	boxHelp.SetTextAlign(tview.AlignCenter)
	boxHelp.SetBackgroundColor(tcell.Color(987))
	helpPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxHelp, 5, 3, false)

	//pagesMain := New       *tview.Pages
	boxTop := tview.NewTextView()
	boxTop.SetBorder(true)
	boxTop.SetText("\n uiRDP - Copyright 2021 Andrew Voytenko")
	boxTop.SetTextAlign(tview.AlignCenter)
	boxTop.SetBackgroundColor(tcell.Color(987))

	boxMidlle := tview.NewBox().SetBorder(true).SetTitle("Midlle")

	//boxBottom
	// boxBottom := tview.NewTextView()
	// boxBottom.SetBorder(true)
	// boxBottom.SetText("\n uiRDP - Copyright 2021 Andrew Voytenko")
	// boxBottom.SetTextAlign(tview.AlignCenter)

	txtMenu := tview.NewTable().SetEvaluateAllRows(true)
	txtMenu.SetBorder(false)
	for i := range ipanelMenu {
		txtMenu.SetCellSimple(0, i, ipanelMenu[i])
		txtMenu.GetCell(0, i).SetAlign(tview.AlignCenter)
	}

	//txtMenu.txAddress = tview.NewTableCell("0.0.0.0:0")
	//txInfo.SetCell(0, 1, infoUI.txAddress)

	mainPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxTop, 5, 3, false).
		AddItem(boxMidlle, 0, 2, true).
		//AddItem(boxBottom, 5, 3, false)
		AddItem(txtMenu, 5, 3, false)

	pages.AddPage("main", mainPage, true, true)
	pages.AddPage("help", helpPage, true, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// Global shortcuts
		switch event.Rune() {
		case 'p':
			//app.SetFocus(boxBottom)
			app.SetFocus(txtMenu)
			return nil
		case 't':
			app.SetFocus(boxTop)
			return nil
		case 'h':
			//app.SetRoot(helpPage, true).EnableMouse(true).Draw()
			pages.SwitchToPage("help")
			//pages.RemovePage("main")

			return nil
		case 'q':
			//app.Stop()
			name, _ := pages.GetFrontPage()
			if name == "help" {
				pages.SwitchToPage("main")
			} else {
				app.Stop()
			}
			return nil
		}

		// // Handle based on current focus. Handlers may modify event
		// switch {
		// case projectPane.HasFocus():
		// 	event = projectPane.handleShortcuts(event)
		// case taskPane.HasFocus():
		// 	event = taskPane.handleShortcuts(event)
		// case taskDetailPane.HasFocus():
		// 	event = taskDetailPane.handleShortcuts(event)
		//}
		return nil

	})

	//app.SetRoot(pages, true)
	app.EnableMouse(true)

	err := app.SetRoot(pages, true)
	pages.SwitchToPage("main")
	app.Run()
	if err != nil {
		panic(err)
	}

}

//[]string{"№", " Group ", " Name ", " IP ", "Port", " Date ", " Comment "},
