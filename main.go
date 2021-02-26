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

	txtMenu := tview.NewTable().SetEvaluateAllRows(true)
	txtMenu.SetBorder(false)
	for i := range ipanelMenu {
		txtMenu.SetCellSimple(0, i, ipanelMenu[i])
		txtMenu.GetCell(0, i).SetAlign(tview.AlignCenter).SetExpansion(5)
	}

	mainPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxTop, 5, 3, false).
		AddItem(boxMidlle, 0, 2, true).
		AddItem(txtMenu, 1, 3, false)

	pages.AddPage("main", mainPage, true, true)
	pages.AddPage("help", helpPage, true, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// Global shortcuts
		switch event.Rune() {
		case 'p':
			//app.SetFocus(boxBottom)
			app.SetFocus(txtMenu)
			//return nil
		case 't':
			app.SetFocus(boxTop)
			//return nil
		case 'h':
			pages.SwitchToPage("help")
			//return nil
		case 'q':
			name, _ := pages.GetFrontPage()
			if name == "help" {
				pages.SwitchToPage("main")
			} else {
				app.Stop()
			}
			//return nil
		}
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
