package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/PyMarcus/gold_watcher/repository"
	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                 fyne.App
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	PriceChartContainer *fyne.Container
	Toolbar             *widget.Toolbar
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	HTTPClient          *http.Client
	Repository          *repository.SQLiteRepository
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferences")
	myApp.App = fyneApp

	myApp.HTTPClient = &http.Client{}

	// log settings
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// database
	sqlDB, err := (func() (*sql.DB, error) {

		path := ""
		if os.Getenv("DB_PATH") != "" {
			path = os.Getenv("DB_PATH")
		} else {
			path = myApp.App.Storage().RootURI().Path() + "sql.db"
		}

		db, err := sql.Open("sqlite", path)

		if err != nil {
			return nil, err
		}
		myApp.InfoLog.Println("Creating database in", path)
		return db, nil
	}())

	if err != nil {
		log.Panic(err)
	}

	myApp.Repository = repository.NewSQLiteRepository(sqlDB)
	err = myApp.Repository.Migrate()

	if err != nil {
		myApp.ErrorLog.Println(err)
		log.Panic()
	}

	// window settings
	myApp.MainWindow = fyneApp.NewWindow("Monitor do Ouro")

	myApp.MainWindow.Resize(fyne.NewSize(WIDTH, HEIGHT))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUi()

	myApp.MainWindow.ShowAndRun()
}
