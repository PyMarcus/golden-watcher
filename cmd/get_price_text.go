package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/PyMarcus/internal/service"
)

func (app *Config) GetPriceText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var open, current, change *canvas.Text

	_gold := service.Gold{}
	gold, err := _gold.GetPrices()

	if err != nil {
		log.Println("Fail to get gold price")
		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}

		open = canvas.NewText("Open Unreachable", grey)
		current = canvas.NewText("Current Unreachable", grey)
		change = canvas.NewText("Close Unreachable", grey)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		if gold.XAUPrice < gold.XAUClose {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
		}

		openTxt := fmt.Sprintf("Open: R$%.4f %s", gold.XAUClose, gold.Currency)
		currentTxt := fmt.Sprintf("Current: R$%.4f %s", gold.XAUPrice, gold.Currency)
		changeTxt := fmt.Sprintf("Change: R$%.4f %s", gold.ChgXAU, gold.Currency)

		open = canvas.NewText(openTxt, nil)
		current = canvas.NewText(currentTxt, displayColor)
		change = canvas.NewText(changeTxt, displayColor)
	}

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return open, current, change
}
