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

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Preços", theme.HomeIcon(), canvas.NewText("Os preços aparecem aqui.", nil)),
		container.NewTabItemWithIcon("Lista", theme.InfoIcon(), canvas.NewText("Preços tabelados.", nil)),
	)

	finalContent := container.NewVBox(priceContainer, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)
}
