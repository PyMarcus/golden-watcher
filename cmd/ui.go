package main

import (
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
}
