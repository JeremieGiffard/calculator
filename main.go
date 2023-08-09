package main

import (
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

	buttonUpdateLabel := widget.NewButton("10/2 * -1", func() {
		HandleClickButton(defaultLabel)

	})
	button1 := CreateButton(2)

	contentContainer := container.New(layout.NewGridLayout(2), defaultLabel, buttonUpdateLabel, button1[0], button1[1])

	myWindow.SetContent(contentContainer)
	myWindow.ShowAndRun()

}

func HandleClickButton(defaultLabel *widget.Label) {
	result, _ := EvaluateCalcul("10/2 * -1")
	defaultLabel.SetText(result)
}

func EvaluateCalcul(expressionToEvaluate string) (string, error) {
	expression, _ := govaluate.NewEvaluableExpression(expressionToEvaluate)
	result, err := expression.Evaluate(nil)
	return fmt.Sprint(result), err
}

func CreateButton(iter int) []*widget.Button {
	var sliceButton []*widget.Button
	item := 0
	for item < iter {
		sliceButton = append(sliceButton, widget.NewButton(fmt.Sprint(item), func() {}))
		item++
	}
	return sliceButton
}
