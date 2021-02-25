package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app   *tview.Application // The tview application.
	pages *tview.Pages       // The application pages.
)

func main() {
	app = tview.NewApplication()
	pages := tview.NewPages()
	// helpPage
	//helpPage := tview.NewFlex()

	boxHelp := tview.NewTextView()
	boxHelp.SetBorder(true)
	boxHelp.SetText("\n[yellow::u] uiRDP - Copyright 2021 Andrew Voytenko [::]")
	boxHelp.SetTextAlign(tview.AlignCenter)
	boxHelp.SetBackgroundColor(tcell.Color(987))
	helpPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxHelp, 5, 3, false)

	//pagesMain := New       *tview.Pages
	boxTop := tview.NewTextView()
	boxTop.SetBorder(true)
	boxTop.SetText("\n[yellow::u] uiRDP - Copyright 2021 Andrew Voytenko [::]")
	boxTop.SetTextAlign(tview.AlignCenter)
	boxTop.SetBackgroundColor(tcell.Color(987))

	boxMidlle := tview.NewBox().SetBorder(true).SetTitle("Midlle")
	//boxBottom := tview.NewBox().SetBorder(true).SetTitle("Bottom")

	boxBottom := tview.NewTextView()
	boxBottom.SetBorder(true)
	boxBottom.SetText("\n uiRDP - Copyright 2021 Andrew Voytenko")
	boxBottom.SetTextAlign(tview.AlignCenter)

	mainPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxTop, 5, 3, false).
		AddItem(boxMidlle, 0, 2, true).
		AddItem(boxBottom, 5, 3, false)

	pages.AddPage("main", mainPage, true, true)
	pages.AddPage("help", helpPage, true, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// Global shortcuts
		switch event.Rune() {
		case 'p':
			app.SetFocus(boxBottom)
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
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

// 2020 uirdp
// загрузить файл конфигурации если его нет то создать
// отсортировать список
// очистить кеш
// получить базовые параметры системы - Время и Дата, Сетевые настройки
// вывести главный экран приложения
//
// ipanelMenu := []string{
// 	"h|Help", "i|Info", "m|Menu", "t|Test", "s|Start", "n|New", "c|Copy", "e|Edit", "d|Dellete", "q|Exit",
// }

//[]string{"№", " Group ", " Name ", " IP ", "Port", " Date ", " Comment "},

// button := tview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
// 	app.Stop()
// })
// button := tview.NewButton("OK").SetSelectedFunc(func() {
// 	app.Stop()
// })
// statusBar = &StatusBar{
// 	Pages:     tview.NewPages(),
// 	message:   tview.NewTextView().SetDynamicColors(true).SetText("Loading..."),
// 	container: app,
// }

// statusBar.AddPage("messagePage", "statusBar.message", true, true)
// boxBottomTxt := tview.NewTextView().
// 	SetDynamicColors(true).
// 	SetText("This is the status bar")

//boxBottom.AddPage(boxBottomTxt)

// flex := tview.NewFlex().
// 	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
// 		AddItem(boxTop, 3, 1, false).
// 		AddItem(boxMidlle, 0, 2, false).
// 		AddItem(boxBottom, 5, 3, false), 0, 2, false)
