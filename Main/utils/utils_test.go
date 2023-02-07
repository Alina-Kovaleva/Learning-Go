package utils

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		a    int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{4, "4"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{17, "17"},
		{45, "FizzBuzz"},
		{99, "Fizz"},
		{55, "Buzz"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.a)
		t.Run(testname, func(t *testing.T) {
			got := FizzBuzz(test.a)
			if got != test.want {
				t.Errorf("FizzBuzz(%d) = %s, wanted %s", test.a, got, test.want)
			}
		})
	}
}

func TestVariables(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want string
	}{
		{4, 4, "EVEN"},
		{5, 5, "ODD"},
		{7, 6, "448"},
		{4, 6, "6.666666666666666E-01"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.a)
		t.Run(testname, func(t *testing.T) {
			got := Variables(test.a, test.b)
			if got != test.want {
				t.Errorf("Variables(%d, %d)  = %s, wanted %s", test.a, test.b, got, test.want)
			}
		})
	}
}

// func TestFizzBuzzString(t *testing.T) {
// 	var tests = []struct {
// 		a    int
// 		want string
// 	}{
// 		{3, "Fizz"},
// 		{5, "Buzz"},
// 		{15, "FizzBuzz"},
// 	}

// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("%d", tt.a)
// 		t.Run(testname, func(t *testing.T) {
// 			got := FizzBuzz(tt.a)
// 			if got != tt.want {
// 				t.Errorf("FizzBuzz(%d) = %s, wanted %s", tt.a, got, tt.want)
// 			}
// 		})
// 	}
// }
