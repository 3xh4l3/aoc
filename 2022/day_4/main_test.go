package main

import (
	"testing"
)

func Test_day_4(t *testing.T) {
	p1, p2 := day_4("example.txt")
	if p1 != 2 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 4 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}
