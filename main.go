package main

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/Knetic/govaluate.v2"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello world")

	defaultLabel := widget.NewLabel("julien")

	buttonUpdateLabel := widget.NewButton("click to Update Label", func() {
		HandleClickButton(defaultLabel)

	})

	contentContainer := container.New(layout.NewGridLayout(2), defaultLabel, buttonUpdateLabel)

	myWindow.SetContent(contentContainer)
	myWindow.ShowAndRun()

}

func HandleClickButton(defaultLabel *widget.Label) {
	expression, _ := govaluate.NewEvaluableExpression("10/2 * -1")
	result, _ := expression.Evaluate(nil)

	defaultLabel.SetText(fmt.Sprint(result))
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
