package main

import (
	"fmt"
	"testing"
)

func Test_day_11(t *testing.T) {
	p1, p2 := day_11("example.txt")
	fmt.Println(p1, p2)
	if p1 != 10605 {
		t.Errorf("Wrong answer in part one %d", p1)
	}
	if p2 != 2713310158 {
		t.Errorf("Wrong answer in part two %d", p2)
	}
}
