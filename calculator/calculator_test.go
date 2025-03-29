package calculator

import (
	"testing"
)

func TestOperations(t *testing.T) {
	tests := []struct {
		name      string
		a         float32
		b         float32
		operator  string
		expected  float32
		wantPanic bool
	}{
		{"Addition", 2, 3, "+", 5, false},
		{"Addition with negative", 2, -2, "+", 0, false},
		{"Addition with negatives", -2, -2, "+", -4, false},
		{"Subtraction", 10, 4, "-", 6, false},
		{"Subtraction with negatives", -10, -4, "-", -6, false},
		{"Subtraction mix", -10, 4, "-", -14, false},
		{"Multiplication", 10, 4, "*", 40, false},
		{"Multiplication with negative", 10, -4, "*", -40, false},
		{"Multiplication with negatives", -10, -4, "*", 40, false},
		{"Division", 10, 2, "/", 5, false},
		{"Division with negative", 10, -2, "/", -5, false},
		{"Division with negatives", -10, -2, "/", 5, false},
		{"Division by zero", 10, 0, "/", 0, true}, // Expected to panic
		{"Power", 2, 2, "^", 4, false},
		{"Power with negative exponent", 2, -2, "^", 0.25, false},
		{"Power with negative base", -2, -2, "^", 0.25, false},
		{"Square Root", 4, 0, "sqrt", 2, false},
		{"Square Root of Negative", -2, 0, "sqrt", 0, true}, // Expected to panic
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Handle expected panic cases
			defer func() {
				if r := recover(); r != nil {
					if !test.wantPanic {
						t.Errorf("Unexpected panic in test: %s, got: %v", test.name, r)
					}
				} else if test.wantPanic {
					t.Errorf("Expected panic but did not occur in test: %s", test.name)
				}
			}()

			if !test.wantPanic {
				result := Calculator(test.a, test.operator, test.b)
				if result != test.expected {
					t.Errorf("%s failed: expected %f, got %f", test.name, test.expected, result)
				}
			} else {
				// If we expect a panic, just call the function and let the defer handle it
				Calculator(test.a, test.operator, test.b)
			}
		})
	}
}
