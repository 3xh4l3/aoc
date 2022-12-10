package main

import (
	"testing"
)

func Test_day_8(t *testing.T) {
	p1, p2 := day_8("example.txt")
	if p1 != 21 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 8 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}
