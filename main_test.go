package main

import (
	"reflect"
	"testing"
)

func TestDivideByZeroShouldThrowError(t *testing.T) {
	_, err := Division(1, 0)
	if err == nil {
		t.Errorf("error %v : divide by 0 should trow error", err)
	}

}
func TestEvaluateCalculShouldReturnTypeString(t *testing.T) {
	got, _ := EvaluateCalcul("1/1")
	got = reflect.TypeOf(got).String()
	want := "string"

	if got != want {
		t.Errorf("got %v type,  want %v type", got, want)
	}
}
