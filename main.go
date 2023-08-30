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

// The makeUI function generate all the UI elements of the program.
func makeUI() (*widget.Label, []*widget.Button) {
	buttonEvaluate := makeButtonEvaluate()
	var sliceSign = []string{"+", "-", "*", "/"}

	var sliceButton []*widget.Button
	sliceButton = append(sliceButton, makeButtonNumber(9)...)
	sliceButton = append(sliceButton, makeSignButton(sliceSign)...)
	sliceButton = append(sliceButton, buttonEvaluate)

	resultLabel = widget.NewLabel("julien")

	return resultLabel, sliceButton

}

// The makeButtonEvaluate function make a equal button with a handler to call the goEvaluate method.
func makeButtonEvaluate() *widget.Button {
	equalButton := widget.NewButton("=", func() {
		log.Println(stringToEvaluate)
		if stringToEvaluate != "" {
			result, error := EvaluateCalcul(stringToEvaluate)
			if error != nil {
				log.Println(error)
				resultLabel.SetText("wrong input")
			} else {
				resultLabel.SetText(result)
			}
			stringToEvaluate = ""
		} else {
			log.Println("Nothing to evaluate")
		}

	})
	return equalButton
}

// The makeButtonNumber take a int arg and make as many button.
func makeButtonNumber(iter int) []*widget.Button {
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

// The makeSignButton take a slice and make button with slice[n] as label
func makeSignButton(sliceSign []string) []*widget.Button {
	var sliceButton []*widget.Button
	for _, v := range sliceSign {
		value := v
		newButton := widget.NewButton(v, func() {
			HandleClickButton(value)
		})
		sliceButton = append(sliceButton, newButton)
	}
	return sliceButton
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello world")
	resultLabel, buttonsUI := makeUI()

	contentContainer := container.New(layout.NewGridLayout(2), resultLabel, buttonsUI[0], buttonsUI[1], buttonsUI[2], buttonsUI[3], buttonsUI[4], buttonsUI[5], buttonsUI[6], buttonsUI[7], buttonsUI[8], buttonsUI[9], buttonsUI[10], buttonsUI[11], buttonsUI[12], buttonsUI[13])

	myWindow.SetContent(contentContainer)
	myWindow.ShowAndRun()

}

// The HandleClickButton is called on click button event.
// Concat a string to another and update the UI with new value
func HandleClickButton(labelButton string) {
	log.Println("item :" + labelButton)
	stringToEvaluate += labelButton
	resultLabel.SetText(stringToEvaluate)
}

// the EvaluateCalcul is a wrapper to manage govaluate return response
func EvaluateCalcul(expressionToEvaluate string) (string, error) {
	var result string
	expression, err := govaluate.NewEvaluableExpression(expressionToEvaluate)
	if err != nil {
		log.Println("error evaluate expression")
		stringToEvaluate = ""
		result = "error evaluate expression"
	} else {
		output, _ := expression.Evaluate(nil)
		result = fmt.Sprint(output)
	}
	return fmt.Sprint(result), err
}
