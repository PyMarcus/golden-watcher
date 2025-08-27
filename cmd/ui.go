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

	finalContent := container.NewVBox(priceContainer)

	app.MainWindow.SetContent(finalContent)
}
