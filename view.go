package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func Graph() *tvxwidgets.UtilModeGauge {
	gauge := tvxwidgets.NewUtilModeGauge()
	gauge.SetLabel("cpu usage:")
	gauge.SetLabelColor(tcell.ColorLightSkyBlue)
	gauge.SetWarnPercentage(65)
	gauge.SetCritPercentage(80)

	return gauge
}

type UIView struct {
	app           *tview.Application
	textView      *tview.TextView
	capGraph      *tvxwidgets.UtilModeGauge
	nodeSparkline *tvxwidgets.Sparkline
	botSparkline  *tvxwidgets.Sparkline
	lineChart     *tvxwidgets.Plot
	barChart      *tvxwidgets.BarChart
	mainWidgets   *tview.Flex
}

func UI(nm *NodeManager) *UIView {
	app := tview.NewApplication()

	textView := tview.NewTextView()
	textView.SetBorder(true).SetTitle("Log")
	textView.SetScrollable(true)
	textView.SetMaxLines(100)
	textView.SetChangedFunc(func() {
		app.Draw()
	})

	capGraph := Graph()

	nodeSparkline := tvxwidgets.NewSparkline()
	nodeSparkline.SetBorder(false)
	nodeSparkline.SetDataTitle("Nodes")
	nodeSparkline.SetDataTitleColor(tcell.ColorOrange)
	nodeSparkline.SetLineColor(tcell.ColorMediumPurple)

	botSparkline := tvxwidgets.NewSparkline()
	botSparkline.SetBorder(false)
	botSparkline.SetDataTitle("Bots")
	botSparkline.SetDataTitleColor(tcell.ColorOrange)
	botSparkline.SetLineColor(tcell.ColorSteelBlue)

	utilFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	utilFlex.AddItem(capGraph, 1, 0, false)
	utilFlex.SetTitle("utilisation mode gauge")
	utilFlex.SetBorder(true)

	sparklineGroupLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	sparklineGroupLayout.SetBorder(false)
	sparklineGroupLayout.SetTitle("Node Stats")
	sparklineGroupLayout.AddItem(nodeSparkline, 0, 1, false)
	sparklineGroupLayout.AddItem(tview.NewBox(), 1, 0, false)
	sparklineGroupLayout.AddItem(botSparkline, 0, 1, false)
	utilFlex.AddItem(sparklineGroupLayout, 0, 1, false)

	lineChart := NewLineChart()
	barChart := newBarChart()

	linebarLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	linebarLayout.SetBorder(false)
	linebarLayout.AddItem(barChart, 0, 1, false)
	linebarLayout.AddItem(tview.NewBox(), 1, 0, false)
	linebarLayout.AddItem(lineChart, 0, 1, false)
	utilFlex.AddItem(linebarLayout, 0, 1, false)

	mainWidgets := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(utilFlex, 0, 1, false).
		AddItem(textView, 0, 1, false)

	modal := tview.NewModal().
		AddButtons([]string{"Start", "CLI"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Start" {
				app.SetRoot(mainWidgets, true).EnableMouse(false)
				nm.NextTick()
				go func() {
					err := app.Run()
					if err != nil {
						panic(err)
					}
					app.Draw()
				}()
			} else if buttonLabel == "CLI" {
				nm.SetInteractive(false)
				app.Stop()
			}
		})

	pages := tview.NewPages().
		AddPage("background", mainWidgets, true, true).
		AddPage("modal", modal, true, true)

	app.SetRoot(pages, true).EnableMouse(true)
	modal.HasFocus()

	return &UIView{
		app:           app,
		capGraph:      capGraph,
		textView:      textView,
		nodeSparkline: nodeSparkline,
		botSparkline:  botSparkline,
		lineChart:     lineChart,
		barChart:      barChart,
		mainWidgets:   mainWidgets,
	}
}

func NewLineChart() *tvxwidgets.Plot {
	bmLineChart := tvxwidgets.NewPlot()
	bmLineChart.SetBorder(false)
	bmLineChart.SetTitle("Capacity")
	bmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorGreen,
		tcell.ColorRed,
	})
	bmLineChart.SetMarker(tvxwidgets.PlotMarkerBraille)

	return bmLineChart
}

func newBarChart() *tvxwidgets.BarChart {
	barGraph := tvxwidgets.NewBarChart()
	barGraph.SetBorder(false)
	barGraph.AddBar("sm", 0, tcell.ColorBlue)
	barGraph.AddBar("med", 0, tcell.ColorLightSkyBlue)
	barGraph.AddBar("lrg", 0, tcell.ColorGreen)
	barGraph.AddBar("busy", 0, tcell.ColorRed)
	barGraph.AddBar("available", 0, tcell.ColorGreen)
	barGraph.AddBar("requests", 0, tcell.ColorRed)
	barGraph.AddBar("avg demand", 0, tcell.ColorMediumVioletRed)
	barGraph.AddBar("total demand", 0, tcell.ColorRed)

	return barGraph
}

func FinishModal(ui *UIView, score string) {
	modal := tview.NewModal().
		AddButtons([]string{"Exit"}).
		SetText(score).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			os.Exit(0)
		})

	pages := tview.NewPages().
		AddPage("background", ui.mainWidgets, true, true).
		AddPage("modal", modal, true, true)

	ui.app.SetRoot(pages, true).Run()
}
