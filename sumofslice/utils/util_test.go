package utils

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{-1, 1},
		{1, 1},
		{0, 0},
	}
	// got := Abs(-1)
	// if got != 1 {
	// 	t.Errorf("Abs(-1) = %d, wanted 1", got)
	// }
	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.a)
		t.Run(testname, func(t *testing.T) {
			got := Abs(test.a)
			if got != test.want {
				t.Errorf("Abs(%d) = %d, wanted %d", test.a, got, test.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{-1, 2, 2},
		{1, 3, 3},
		{0, 4, 4},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.a)
		t.Run(testname, func(t *testing.T) {
			got := Max(test.a, test.b)
			if got != test.want {
				t.Errorf("Max(%d, %d) = %d, wanted %d", test.a, test.b, got, test.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{-1, 2, -1},
		{1, 3, 1},
		{0, 4, 0},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.a)
		t.Run(testname, func(t *testing.T) {
			got := Min(test.a, test.b)
			if got != test.want {
				t.Errorf("Max(%d, %d) = %d, wanted %d", test.a, test.b, got, test.want)
			}
		})
	}
}
