package main

import (
	"errors"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello world")

	w.SetContent(widget.NewLabel(hello("julien")))
	w.ShowAndRun()
}

func hello(name string) string {
	return "Hi, " + name
}

type Fraction struct {
	nominator   float64
	denominator float64
}

func (frac Fraction) decimal() float64 {
	return frac.nominator / frac.denominator
}

func Division(fraction1, fraction2 float64) (float64, error) {
	if fraction2 != 0 {
		return fraction1 / fraction2, nil
	} else {
		return 0, errors.New("tried to divide by 0")
	}

}
