package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// Get the current price gold
	openPrice, currentPrice, priceChange := app.getPriceText()

	// Put price information into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	// Get toolbar
	toolBar := app.getToolBar()
	app.Toolbar = toolBar

	priceTabContent := app.pricesTab()

	// Get app tabs

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.HomeIcon(), canvas.NewText("Holdings content goes to here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Add container to window
	finalContent := container.NewVBox(priceContent, toolBar, tabs)

	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 30) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Print("refreshing prices")
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()

}
