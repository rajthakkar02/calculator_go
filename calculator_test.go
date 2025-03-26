package main

import (
	"testing"
)

func TestOperations(t *testing.T) {
	tests := []struct {
		name      string
		function  func(float32, string, ...float32) float32
		a         float32
		b         string
		c         float32
		expected  float32
		wantPanic bool
	}{
		{"Addition", Calculator, 2, "+", 3, 5, false},
		{"Addition", Calculator, 2, "+", -2, 0, false},
		{"Addition", Calculator, -2, "+", -2, -4, false},
		{"Subtraction", Calculator, 10, "-", 4, 6, false},
		{"Subtraction", Calculator, -10, "-", -4, -6, false},
		{"Subtraction", Calculator, -10, "-", 4, -14, false},
		{"Multiplication", Calculator, 10, "*", 4, 40, false},
		{"Multiplication", Calculator, 10, "*", -4, -40, false},
		{"Multiplication", Calculator, -10, "*", -4, 40, false},
		{"Diviosn", Calculator, 10, "/", 2, 5, false},
		{"Diviosn", Calculator, 10, "/", -2, -5, false},
		{"Diviosn", Calculator, -10, "/", -2, 5, false},
		{"Power", Calculator, 2, "^", 2, 4, false},
		{"Power", Calculator, 2, "^", -2, 0.25, false},
		{"Power", Calculator, -2, "^", -2, 0.25, false},
		{"Power of two", Calculator, 2, "^2", 0, 4, false},
		{"Power of two", Calculator, -2, "^2", 0, 4, false},
		{"Power of cube", Calculator, -2, "^3", 0, -8, false},
		{"Power of cube", Calculator, -2, "^3", 0, -8, false},
		{"Power of cube", Calculator, 2, "^3", 0, 8, false},
		{"Square Root", Calculator, 4, "sqrt", 0, 2, false},
		{"Square of Root", Calculator, -2, "sqrt", 0, 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.function(test.a, test.b, test.c)
			if result != test.expected {
				t.Errorf("%s(%f, %s, %f) = %f; want %f", test.name, test.a, test.b, test.c, result, test.expected)
			}
		})
	}
}
