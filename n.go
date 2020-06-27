package main

import (
	"io/ioutil"
	"log"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"

)

func main() {
	app := app.New()
	input, err := ioutil.ReadFile("/dev/stdin")
	if err != nil {
		log.Fatal(err)
	}
	w := app.NewWindow("n")
	label := widget.NewLabel(string(input))
	label.Wrapping = fyne.TextWrapBreak
	scroll := widget.NewScrollContainer(label)
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if(ev.Name == "Q") {
			app.Quit()
		}
		})
	w.SetContent(scroll)
	w.ShowAndRun()

}
