package main

import (
	"testing"
)

func Test_day_6(t *testing.T) {
	p1, p2 := day_6("example.txt")
	if p1 != 11 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 26 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}

func Test_day_6_1(t *testing.T) {
	p1, p2 := day_6("example1.txt")
	if p1 != 10 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 29 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}

func Test_day_6_2(t *testing.T) {
	p1, p2 := day_6("example2.txt")
	if p1 != 6 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 23 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}
