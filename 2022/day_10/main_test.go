package main

import (
	"testing"
)

func Test_day_10(t *testing.T) {
	p1, _ := day_10("example.txt")
	if p1 != 13140 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
}
