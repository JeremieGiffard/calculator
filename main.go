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

type CustomButton struct {
	Label   string
	Handler func()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello world")
	resultLabel, buttonsUI := makeUI()
	//    TODO : makeUI should return *fyne.Container or *container.AppTabs
	// Don't switch to fyne.CanvasObject. Make testing not fun at all
	contentContainer := container.New(layout.NewGridLayout(3), resultLabel, layout.NewSpacer(), buttonsUI[0], buttonsUI[1], buttonsUI[2], buttonsUI[3], buttonsUI[4], buttonsUI[5], buttonsUI[6], buttonsUI[7], buttonsUI[8], buttonsUI[9], buttonsUI[10], buttonsUI[11], buttonsUI[12], buttonsUI[13], buttonsUI[14])

	tabs := container.NewAppTabs(
		container.NewTabItem("Basic", contentContainer),
		container.NewTabItem("Currency", widget.NewLabel("World!")),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

// The makeUI function generate all the UI elements of the program.
func makeUI() (*widget.Label, []*widget.Button) {
	var sliceSign = []string{"+", "-", "*", "/"}
	var sliceNumber = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	listNumberButtonStruct := makeSimpleCalculButton(sliceNumber)
	listSignButtonStruct := makeSimpleCalculButton(sliceSign)
	buttonEvaluate := makeButtonEvaluate()

	var sliceButton []*widget.Button
	sliceButton = append(sliceButton, makeButtonsFromList(listNumberButtonStruct)...)
	sliceButton = append(sliceButton, makeButtonsFromList(listSignButtonStruct)...)
	sliceButton = append(sliceButton, buttonEvaluate)

	resultLabel = widget.NewLabel("start")

	return resultLabel, sliceButton

}

// The makeButtonEvaluate function make a equal button with a handler to call the goEvaluate method.
func makeButtonEvaluate() *widget.Button {
	equalButton := widget.NewButton("=", func() {
		log.Println(stringToEvaluate)
		if stringToEvaluate != "" {
			result, err := EvaluateCalcul(stringToEvaluate)
			if err != nil {
				log.Println(err)
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

// The makeCustomButton take a slice and make button with slice[n] as label
func makeSimpleCalculButton(sliceSign []string) []*CustomButton {
	var sliceCustomButton []*CustomButton
	for _, v := range sliceSign {
		item := v
		sliceCustomButton = append(sliceCustomButton, makeCustomStruct(item, func() { HandleClickButton(item) }))
	}
	return sliceCustomButton
}
func makeCustomStruct(arg string, argFunc func()) *CustomButton {
	return &CustomButton{arg, argFunc}
}

func makeButtonsFromList(makeMapStringFunc []*CustomButton) []*widget.Button {
	var sliceButton []*widget.Button
	for _, value := range makeMapStringFunc {
		newButton := widget.NewButton(value.Label, value.Handler)
		sliceButton = append(sliceButton, newButton)
	}
	return sliceButton
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
