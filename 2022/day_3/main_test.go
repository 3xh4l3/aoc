package main

import (
	"testing"
)

func Test_day_3(t *testing.T) {
	sum1, sum2 := day_3("example.txt")
	if sum1 != 157 {
		t.Errorf("Wrong answer in part one %d", sum1)
	}
	if sum2 != 70 {
		t.Errorf("Wrong answer in part two %d", sum2)
	}
}
