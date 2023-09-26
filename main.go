package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/Knetic/govaluate.v2"
)

var (
	stringToEvaluate    = ""
	resultLabel         *widget.Label
	inputCurrency       = widget.NewEntry()
	CurrencyEndPointURL = "https://duckduckgo.com/js/spice/currency/_PARAM_/eur/usd"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello world")
	containerUI := makeUI()

	myWindow.SetContent(containerUI)
	myWindow.ShowAndRun()
}

// The makeUI function generate all the UI elements of the program.
func makeUI() *container.AppTabs {
	var sliceSign = []string{"+", "-", "*", "/"}
	var sliceNumber = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	buttonEvaluate := makeButtonEvaluate()

	var sliceButton []*widget.Button
	sliceButton = append(sliceButton, makeSimpleCalculButton(sliceNumber)...)
	sliceButton = append(sliceButton, makeSimpleCalculButton(sliceSign)...)
	sliceButton = append(sliceButton, buttonEvaluate)

	resultLabel = widget.NewLabel("start")

	contentContainer := container.New(layout.NewGridLayout(3), resultLabel, layout.NewSpacer(), sliceButton[0], sliceButton[1], sliceButton[2], sliceButton[3], sliceButton[4], sliceButton[5], sliceButton[6], sliceButton[7], sliceButton[8], sliceButton[9], sliceButton[10], sliceButton[11], sliceButton[12], sliceButton[13], sliceButton[14])

	labelCurrency := widget.NewLabel("World!")
	inputCurrency.SetPlaceHolder("Enter currency...")

	convertCurrencyBtn := widget.NewButton("convert", func() {

		res, _ := HttpConnect(strings.Replace(CurrencyEndPointURL, "_PARAM_", inputCurrency.Text, 1))
		log.Println(res)
		labelCurrency.SetText(res)

	})
	currencyContainer := container.New(layout.NewGridLayout(1), labelCurrency, inputCurrency, convertCurrencyBtn)

	tabs := container.NewAppTabs(
		container.NewTabItem("Basic", contentContainer),
		container.NewTabItem("Currency", currencyContainer),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs

}

// The makeButtonEvaluate function make a equal button with a handler to call the goEvaluate method.
func makeButtonEvaluate() *widget.Button {
	equalButton := widget.NewButton("=", handleEqualButton())
	return equalButton
}

// The makeCustomButton take a slice and make button with slice[n] as label
func makeSimpleCalculButton(sliceSign []string) []*widget.Button {
	var sliceButton []*widget.Button
	for _, value := range sliceSign {
		sliceButton = append(sliceButton, widget.NewButton(value, HandleClickButton(value)))
	}
	return sliceButton
}

// the handleEqualButton is called on equal button click event.
// Handle error logic for EvaluateCalcul return value
func handleEqualButton() func() {
	return func() {
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
	}
}

// The HandleClickButton is called on click button event.
// Concat a string to another and update the UI with new value
func HandleClickButton(labelButton string) func() {
	return func() {
		log.Println("item :" + labelButton)
		stringToEvaluate += labelButton
		resultLabel.SetText(stringToEvaluate)
	}
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

func HttpConnect(stringUrl string) (string, error) {
	var sb string
	resp, err := http.Get(stringUrl)
	if err != nil {
		log.Fatalln(err)
		sb = "error get request"
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb = string(body)
	// log.Printf("## in ifElse :" + sb)

	log.Printf("## before return :" + sb)
	return sb, err
}
