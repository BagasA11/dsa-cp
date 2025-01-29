package swap

import "testing"

func TestSwap(t *testing.T) {
	a, b := 1, 2
	c, d := a, b
	Swap(&a, &b)
	if a == c || b == d {
		t.Error("gagal swap")
	}
}
