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
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	Toolbar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HTTPClient          *http.Client
}

var myApp Config

func main() {

	// Create a fyne application
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// Create our looger
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open a connection to the database

	// Create a database repository

	// Create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// Show and run the applicaton
	myApp.MainWindow.ShowAndRun()
}
