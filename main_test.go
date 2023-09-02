package main

import (
	"reflect"
	"testing"
)

func TestEvaluateCalculShouldReturnTypeString(t *testing.T) {
	got, _ := EvaluateCalcul("1/1")
	got = reflect.TypeOf(got).String()
	want := "string"

	if got != want {
		t.Errorf("got %v type,  want %v type", got, want)
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
