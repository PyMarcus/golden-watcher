package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PyMarcus/gold_watcher/repository"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.HoldingsTable = app.getHoldingTable()

	_container := container.NewBorder(nil, nil, nil, nil, container.NewAdaptiveGrid(1, app.HoldingsTable))
	return _container
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Quantidade", "Preço de compra", "Comprou em", "Apagar?"})

	for _, x := range holdings {
		var currentRow []any
		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("R$ %d toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("R$ %.2f", float32(x.PurchasePrice)))
		currentRow = append(currentRow, x.PurchaseDate.Local().UTC().Format("02/01/2006 15:04:05"))
		currentRow = append(currentRow, widget.NewButton("Remover", func() {}))
		slice = append(slice, currentRow)
	}

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
	data := app.getHoldingSlice()
	app.Holdings = data

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			_container := container.NewVBox(widget.NewLabel(""))
			return _container
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(data[0])-1) && i.Row != 0 {
				w := widget.NewButtonWithIcon("Remover", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Remover?", "", func(deleted bool) {
						id, _ := strconv.Atoi(data[i.Row][0].(string))
						err := app.Repository.DeleteHolding(int64(id))

						if err != nil {
							app.ErrorLog.Println(err)
						}

						app.refreshTable()

					}, app.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(data[i.Row][i.Col].(string)),
				}
			}
		})

	colWidth := []float32{50, 200, 200, 200}
	for i := 0; i < len(colWidth); i++ {
		table.SetColumnWidth(i, colWidth[i])
	}

	return table
}
