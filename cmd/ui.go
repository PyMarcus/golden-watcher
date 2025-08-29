package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUi() {
	openPrice, currentPrice, changePrice := app.GetPriceText()

	priceContainer := container.NewGridWithColumns(
		3,
		openPrice,
		currentPrice,
		changePrice,
	)
	app.PriceContainer = priceContainer

	toolbar := app.getToolBar()
	app.Toolbar = toolbar

	priceTabContente := app.pricesTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Preços", theme.HomeIcon(), priceTabContente),
		container.NewTabItemWithIcon("Lista", theme.InfoIcon(), canvas.NewText("Preços tabelados.", nil)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	priceChart := app.pricesTab()
	app.PriceChartContainer = priceChart

	finalContent := container.NewVBox(priceContainer, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 5) {
			app.refreshPrices()
		}
	}()

}

func (app *Config) refreshPrices() {
	app.InfoLog.Print("Refreshing prices...")
	open, current, change := app.GetPriceText()
	chart := app.getChart()

	fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
		app.PriceContainer.Refresh()

		app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
		app.PriceChartContainer.Refresh()
	}, false)

}
