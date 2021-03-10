package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func getMainPage() *tview.Flex {
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

	return mainPage

}

//[]string{"№", " Group ", " Name ", " IP ", "Port", " Date ", " Comment "},
