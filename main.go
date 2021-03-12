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
	getConfig()
	app = tview.NewApplication()
	pages := tview.NewPages()

	pages.AddPage("main", getMainPage(), true, true)
	pages.AddPage("help", getHelpPage(), true, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		nameItPage, _ := pages.GetFrontPage()
		// Global shortcuts
		switch event.Key() {
		case tcell.KeyESC:
			if nameItPage == "help" {
				pages.SwitchToPage("main")
			}
		}

		switch event.Rune() {

		case 'h':
			pages.SwitchToPage("help")
		case 'i':

		case 'm':

		case 'q':
			name, _ := pages.GetFrontPage()
			if name == "help" {
				pages.SwitchToPage("main")
				//pages.Draw(tcell.Screen)
				//pages.Blur()
				//app.SetFocus(mainBodyTable)
			} else {
				app.Stop()
			}
		}
		return event
	})

	pages.SwitchToPage("main")
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
