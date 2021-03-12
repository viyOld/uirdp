package main

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	headTable = []string{
		"№  ",
		"Group               ",
		"Name                     ",
		"IP              ",
		"Port   ",
		"Date        ",
		"Comment"}
)

func getMainPage() *tview.Flex {
	boxTop := tview.NewTextView()
	boxTop.SetBorder(true)
	boxTop.SetText("\n uiRDP - Copyright 2021 Andrew Voytenko")
	boxTop.SetTextAlign(tview.AlignCenter)
	boxTop.SetBackgroundColor(tcell.Color(987))

	//boxMidlle := tview.NewBox().SetBorder(true).SetTitle("Midlle")
	mainBodyTable := tview.NewTable().SetEvaluateAllRows(true).SetFixed(1, 10)
	mainBodyTable.SetBorder(true).SetBorderPadding(2, 2, 3, 3)
	mainBodyTable.SetSelectable(true, false)
	//mainBodyTable.SetColumnWidths([]int{4, 25, 25, 17, 7, 12, 0})
	//mainBodyTable.SetRect()
	//mainBodyTable.SetBorderPadding(2, 2, 3, 3)
	for i := range headTable {
		mainBodyTable.SetCellSimple(0, i, headTable[i])
		mainBodyTable.GetCell(0, i).SetAlign(tview.AlignLeft).SetSelectable(false).SetTextColor(tcell.ColorOrange) //.SetExpansion(1)
		//cell := mainBodyTable.GetCell(0, i)
	}
	for j := range config.Servers {
		mainBodyTable.SetCellSimple(j+1, 0, strconv.Itoa(j+1)) //SetMaxWidth(3)
		mainBodyTable.SetCellSimple(j+1, 1, config.Servers[j].Group)
		mainBodyTable.SetCellSimple(j+1, 2, config.Servers[j].Name)
		mainBodyTable.SetCellSimple(j+1, 3, config.Servers[j].IP)
		mainBodyTable.SetCellSimple(j+1, 4, strconv.Itoa(config.Servers[j].Port))
		mainBodyTable.SetCellSimple(j+1, 5, config.Servers[j].Data)
		mainBodyTable.SetCellSimple(j+1, 6, config.Servers[j].Comment)
	}

	txtMenu := tview.NewTable().SetEvaluateAllRows(true)
	txtMenu.SetBorder(false)
	for i := range ipanelMenu {
		txtMenu.SetCellSimple(0, i, ipanelMenu[i])
		txtMenu.GetCell(0, i).SetAlign(tview.AlignCenter).SetExpansion(5)
	}

	mainPage := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(boxTop, 5, 3, false).
		AddItem(mainBodyTable, 0, 2, true).
		AddItem(txtMenu, 1, 3, false)

	return mainPage

}
