package main

import (
	"testing"
)

func Test_day_9(t *testing.T) {
	p1, p2 := day_9("example.txt")
	if p1 != 13 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 1 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}
