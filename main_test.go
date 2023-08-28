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
