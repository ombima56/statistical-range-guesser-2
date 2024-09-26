package calculation_test

import (
	"math"
	"testing"

	"learn.zone01kisumu.ke/git/hiombima/guess-it-2/calculation"
)

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		data []float64
		expectedSlope float64
		expectedIntercept float64
	} {
		{
			data: []float64{1,2,3,4,5},
			expectedSlope: 1.0,
			expectedIntercept: 1.0,
		},
		{
			data:              []float64{10, 12, 14, 16, 18},
			expectedSlope:     2.0,
			expectedIntercept: 10.0,
		},
	}

	for _, test := range tests {
		slope, intercept := calculation.LinearRegression(test.data)

		if math.Abs(slope-test.expectedSlope) > 1e-9 {
			t.Errorf("for data %v, expected slope %.2f but got %.2f", test.data, test.expectedSlope, slope)
		}

		if math.Abs(intercept-test.expectedIntercept) > 1e-9 {
			t.Errorf("For data %v, expected intercept %.2f but got %.2f", test.data, test.expectedIntercept, intercept)
		}
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		data []float64
		expected float64
	}{
		{
			name: "Average of multiple positive numbers",
			data: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: 3.0,
		},
		{
			name: "Average of multiple negative numbers",
			data: []float64{-1.0, -2.0, -3.0, -4.0, -5.0},
			expected: -3.0,
		},
		{
			name: "Average of a mix of positive and negative numbers",
			data: []float64{-1.0, 1.0, -1.0, 1.0},
			expected: 0.0,
		},
		{
			name:     "Average of a single number",
			data:     []float64{10.0},
			expected: 10.0,
		},
		{
			name:     "Average of empty data",
			data:     []float64{},
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculation.Average(test.data)
			if result != test.expected {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}