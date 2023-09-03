package main

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2/test"
)

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

func TestMakeSimpleCalculButton(t *testing.T) {
	var sliceNumber = []string{"0", "1", "2", "3", "4", "5"}
	got := makeSimpleCalculButton(sliceNumber)
	want := 6

	if len(got) != want {
		t.Errorf("got %v length, want %v ", len(got), want)
	}
	if got[1].Label != sliceNumber[1] {
		t.Errorf("got %v label, want %v ", got[1].Label, sliceNumber[1])
	}
}

func TestMakeUILabel(t *testing.T) {
	labelUI, _ := makeUI()

	got := reflect.TypeOf(labelUI).String()
	want := "*widget.Label"
	if got != want {
		t.Errorf("got %v type, want %v ", got, want)
	}
}

func TestMakeUIButton(t *testing.T) {
	_, sliceButton := makeUI()

	got := sliceButton[1].Text
	want := "1"
	if got != want {
		t.Errorf("got text buttin %v , want %v ", got, want)
	}
}

func TestHandleClickButton(t *testing.T) {
	labelUI, _ := makeUI()
	HandleClickButton("1")

	got := labelUI.Text
	want := "1"
	if got != want {
		t.Errorf("got text label %v , want %v ", got, want)
	}
}

func TestMakeButtonEvaluate(t *testing.T) {
	stringToEvaluate = ""
	labelUI, SliceButtons := makeUI()
	HandleClickButton("1+4")

	test.Tap(SliceButtons[14])
	got := labelUI.Text
	want := "5"
	if got != want {
		t.Errorf("got text label %v , want %v ", got, want)
	}

}

func TestMakeButtonEvaluateErr(t *testing.T) {
	//reset global var
	stringToEvaluate = ""
	labelUI, SliceButtons := makeUI()
	HandleClickButton("+/*")

	test.Tap(SliceButtons[14])

	got := labelUI.Text
	want := "wrong input"
	if got != want {
		t.Errorf("got text label %v , want %v ", got, want)
	}
	//reset global var
	stringToEvaluate = ""
}
func TestMakeButtonEvaluateEmpty(t *testing.T) {
	//reset global var
	stringToEvaluate = ""
	labelUI, SliceButtons := makeUI()
	want := labelUI.Text
	test.Tap(SliceButtons[14])

	got := labelUI.Text
	if got != want {
		t.Errorf("Text label should not be updated. got  %v , want %v ", got, want)
	}
	//reset global var
	stringToEvaluate = ""
}
