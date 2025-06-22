package helpers

import (
	"fmt"
	"testing"
)

func TestX111(t *testing.T) {
	var tests = []struct {
		x, y int
	}{
		{0, 1},
		{1, 3},
		{2, 7},
		{3, 13},
		{8, 73},
		{10, 111},
		{20, 421},
		{30, 931},
		{40, 1641},
		{100, 10101},
		{2000, 4002001},
		{500000, 250000500001},
	}

	for index, tUnit := range tests {
		tName := fmt.Sprintf("test %d: val %d", index, tUnit.x)
		t.Run(tName, func(t *testing.T) {
			res := X111(tUnit.x)
			if tUnit.y != res {
				t.Errorf("wanted: %d, got: %d", tUnit.y, res)
			}
		})
	}
}

func TestDivP(t *testing.T) {
	var tests = []struct {
		n   int
		p   int
		res bool
	}{
		{1, 1, true},
		{2, 1, true},
		{2, 2, true},
		{4, 2, true},
		{4, 3, false},
		{5, 2, false},
		{7, 1, true},
		{7, 2, false},
		{10, 2, true},
		{10, 5, true},
		{11, 3, false},
	}

	for _, tUnit := range tests {
		tName := fmt.Sprintf("values_%d_%d", tUnit.n, tUnit.p)
		t.Run(tName, func(t *testing.T) {
			res := DivP(tUnit.n, tUnit.p)
			if res != tUnit.res {
				t.Errorf("wanted: %t, got: %t", tUnit.res, res)
			}
		})
	}
}
