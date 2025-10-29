package main

import "testing"

var tests = []struct{
	testName string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"happyPath", 100.0, 10.0, 10.0, false},
	{"failurePath", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range(tests) {
		got, error := divide(tt.dividend, tt.divisor)
		if tt.isErr && error == nil {
			t.Error("Error expected but not thrown")
		} else if !tt.isErr && error != nil {
			t.Error("Error not expected but thrown")
		}
		if got != tt.expected {
			t.Errorf("Expected = %v not same as got %v", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Error thrown when it should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Error not thrown when it should have")
	}
}
