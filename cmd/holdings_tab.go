package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/PyMarcus/gold_watcher/repository"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getHoldingSlice() [][]any {
	var slice [][]any

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.Repository.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}
	return holdings, nil
}

func (app *Config) getHoldingTable() *widget.Table {
	return nil
}
