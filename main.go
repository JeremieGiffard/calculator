package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/Knetic/govaluate.v2"
)

var stringToEvaluate = ""
var resultLabel *widget.Label

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello world")

	resultLabel = widget.NewLabel("julien")

	buttonPlus := widget.NewButton("+", func() {
		HandleClickButton("+")
	})
	buttonMinus := widget.NewButton("-", func() {
		HandleClickButton("-")
	})
	buttonDivide := widget.NewButton("/", func() {
		HandleClickButton("/")
	})
	buttonMultiple := widget.NewButton("*", func() {
		HandleClickButton("*")
	})

	buttonEqual := widget.NewButton("=", func() {
		log.Println(stringToEvaluate)
		if stringToEvaluate != "" {
			result, _ := EvaluateCalcul(stringToEvaluate)
			resultLabel.SetText(result)
			stringToEvaluate = ""
		} else {
			log.Println("Nothing to evaluate")
		}

	})
	buttonsNumber := CreateButton(9)

	contentContainer := container.New(layout.NewGridLayout(2), resultLabel, buttonsNumber[0], buttonsNumber[1], buttonsNumber[2], buttonsNumber[3], buttonsNumber[4], buttonsNumber[5], buttonsNumber[6], buttonsNumber[7], buttonsNumber[8], buttonPlus, buttonMinus, buttonDivide, buttonMultiple, buttonEqual)

	myWindow.SetContent(contentContainer)
	myWindow.ShowAndRun()

}

func CreateButton(iter int) []*widget.Button {
	var sliceButton []*widget.Button
	item := 0
	for item < iter {
		value := fmt.Sprint(item)
		newButton := widget.NewButton(fmt.Sprint(item), func() {
			HandleClickButton(value)
		})
		item++
		sliceButton = append(sliceButton, newButton)

	}
	return sliceButton
}
func HandleClickButton(labelButton string) {
	log.Println("item :" + labelButton)
	stringToEvaluate += labelButton
	resultLabel.SetText(stringToEvaluate)
}

func EvaluateCalcul(expressionToEvaluate string) (string, error) {
	expression, _ := govaluate.NewEvaluableExpression(expressionToEvaluate)
	result, err := expression.Evaluate(nil)
	return fmt.Sprint(result), err
}
