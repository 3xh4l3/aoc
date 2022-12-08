package main

import (
	"testing"
)

func Test_day_7(t *testing.T) {
	p1, p2 := day_7("example.txt")
	if p1 != 95437 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 24933642 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}
