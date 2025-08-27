package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App            fyne.App
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	Toolbar        *widget.Toolbar
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	HTTPClient     *http.Client
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferences")
	myApp.App = fyneApp

	myApp.HTTPClient = &http.Client{}

	// log settings
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// window settings
	myApp.MainWindow = fyneApp.NewWindow("Monitor do Ouro")

	myApp.MainWindow.Resize(fyne.NewSize(WIDTH, HEIGHT))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUi()

	myApp.MainWindow.ShowAndRun()
}
