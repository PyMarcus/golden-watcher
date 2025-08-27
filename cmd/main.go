package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/PyMarcus/settings"
)

type Config struct {
	App            fyne.App
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferences")
	myApp.App = fyneApp

	// log settings
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// window settings
	myApp.MainWindow = fyneApp.NewWindow("Monitor do Ouro")

	myApp.MainWindow.Resize(fyne.NewSize(settings.WIDTH, settings.HEIGHT))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUi()

	myApp.MainWindow.ShowAndRun()
}
