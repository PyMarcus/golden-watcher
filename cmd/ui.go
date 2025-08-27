package main

import "fyne.io/fyne/v2/container"

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

	finalContent := container.NewVBox(priceContainer, toolbar)

	app.MainWindow.SetContent(finalContent)
}
