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

func TestCreateButton(t *testing.T) {
	got := makeButtonNumber(5)
	want := 5

	if len(got) != want {
		t.Errorf("got %v length, want %v ", len(got), want)
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
