package main

import (
	"goldwatcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getHoldingsTable() *widget.Table {
	return nil
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	var h []repository.Holdings

	return h, nil
}