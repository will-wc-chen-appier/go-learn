package main

import (
	"fmt"
	"testing"
)

// return the smaller int
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, 3)
	if ans != 2 {
		t.Errorf("IntMin(2,3) = %d; want 2", ans) //reports an error
		//t.Fatalf("IntMin(2,3) = %d; want 2", ans) //reports an error, and stops the whole test
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 2, 1},
		{3, 0, 0},
		{-1, -1, -1},
		{100, 50, 50},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%d_%d", test.a, test.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(test.a, test.b)
			if ans != test.want {
				t.Errorf("IntMin(%d,%d) = %d; want %d", test.a, test.b, ans, test.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		IntMin(1, 2)
	}
}
