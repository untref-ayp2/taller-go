package main

import "testing"

func TestSwap(t *testing.T) {
	x, y := 10, 20
	Swap(&x, &y)
	if x != 20 || y != 10 {
		t.Errorf("Swap(10,20) = (%d,%d); esperado (20,10)", x, y)
	}
}
