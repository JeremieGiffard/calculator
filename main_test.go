package main

import (
	"reflect"
	"strings"
	"testing"

	"fyne.io/fyne/v2/test"
)

func Cleanup() {
	//reset global var
	stringToEvaluate = ""
}

func TestEvaluateCalculShouldReturnTypeString(t *testing.T) {
	got, _ := EvaluateCalcul("1/1")
	got = reflect.TypeOf(got).String()
	want := "string"

	if got != want {
		t.Errorf("got %v type,  want %v type", got, want)
	}
}

func TestEvaluateCalculErr(t *testing.T) {
	got, errors := EvaluateCalcul("-+/")

	if errors == nil {
		t.Errorf("got nil value , wanted a error")
	}

	want := "error evaluate expression"
	if got != want {
		t.Errorf("got %v , want %v ", got, want)
	}

}

func TestMakeUILabel(t *testing.T) {
	t.Cleanup(Cleanup)

	containerUI := makeUI()
	got := len(containerUI.Items)

	want := 2
	if got != want {
		t.Errorf("got %v items, want %v ", got, want)
	}
}

func TestMakeSimpleCalculButton(t *testing.T) {
	t.Cleanup(Cleanup)
	var sliceNumber = []string{"0", "1", "2", "3", "4", "5"}
	got := makeSimpleCalculButton(sliceNumber)
	want := 6

	if len(got) != want {
		t.Errorf("got %v length, want %v ", len(got), want)
	}
	if got[1].Text != sliceNumber[1] {
		t.Errorf("got %v label, want %v ", got[1].Text, sliceNumber[1])
	}
}

func TestMakeUIButton(t *testing.T) {
	tabsContainer := makeUI()

	got := tabsContainer.Items[0].Text
	want := "Basic"
	if got != want {
		t.Errorf("got tab label %v , want %v ", got, want)
	}
}

func TestHandleClickButton(t *testing.T) {
	t.Cleanup(Cleanup)
	//test button handler HandleClickButton by Tap on a test button
	test.NewApp()
	makeUI()

	var sliceNumber = []string{"0", "1", "2"}
	testclickValue := makeSimpleCalculButton(sliceNumber)
	test.Tap(testclickValue[1])

	got := resultLabel.Text
	want := "1"
	if got != want {
		t.Errorf("got text label %v , want %v ", got, want)
	}
}

func TestMakeButtonEvaluate(t *testing.T) {
	t.Cleanup(Cleanup)
	test.NewApp()
	makeUI()

	buttonEvaluate := makeButtonEvaluate()

	var sliceNumber = []string{"1", "+", "2"}
	testClickValue := makeSimpleCalculButton(sliceNumber)

	test.Tap(testClickValue[0]) //"1"
	test.Tap(testClickValue[1]) //"1+"
	test.Tap(testClickValue[2]) // "1+2"
	test.Tap(buttonEvaluate)    //should handle"3"

	got := resultLabel.Text
	want := "3"
	if got != want {
		t.Errorf("got text label %v , want %v ", got, want)
	}

}

func TestMakeButtonEvaluateErr(t *testing.T) {
	t.Cleanup(Cleanup)
	makeUI()
	buttonEvaluate := makeButtonEvaluate()

	var sliceNumber = []string{"+", "/", "*"}
	testclickValue := makeSimpleCalculButton(sliceNumber)
	test.Tap(testclickValue[0]) //"+"
	test.Tap(testclickValue[1]) //"+/"
	test.Tap(testclickValue[2]) // "+/*"
	test.Tap(buttonEvaluate)

	got := resultLabel.Text

	want := "wrong input"
	if got != want {
		t.Errorf("got text label %v , want %v ", got, want)
	}
	//reset global var
	stringToEvaluate = ""
}
func TestMakeButtonEvaluateEmpty(t *testing.T) {
	makeUI()
	want := resultLabel.Text
	handleEqualButton()
	got := resultLabel.Text

	if got != want {
		t.Errorf("Text label should not be updated. got  %v , want %v ", got, want)
	}
	//reset global var
	stringToEvaluate = ""
}

func TestHttpConnectCurrentEndPoint(t *testing.T) {

	resp, err := HttpConnect("https://duckduckgo.com/js/spice/currency/12/eur/usd")

	if err != nil {
		t.Errorf("err should be nil.Instead got  %v ", err)
	}
	want := "ddg_spice_currency"
	if !strings.Contains(resp, want) {
		t.Errorf("body should not be empty .Instead got  %v ", resp)
	}
}
