package main

import (
	"testing"
)

func Test_day_5(t *testing.T) {
	p1, p2 := day_5("example.txt")
	if p1 != "CMZ" {
		t.Errorf("Wrong answer in part one %s", p1)
	}
	if p2 != "MCD" {
		t.Errorf("Wrong answer in part two %s", p2)
	}
}
