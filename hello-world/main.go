package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppLabel struct {
	output *widget.Label
}

var myAppLabel AppLabel

func main() {

	a := app.New()
	w := a.NewWindow("Hello World!")

	output, entry, btn := myAppLabel.makeUI()

	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()

}

func (app *AppLabel) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {

	output := widget.NewLabel("Hello world!")

	entry := widget.NewEntry()

	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	app.output = output

	return output, entry, btn
}
