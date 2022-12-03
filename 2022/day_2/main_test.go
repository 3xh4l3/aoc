package main

import (
	"testing"
)

func Test_day_2(t *testing.T) {
	s1, s2 := day_2("example.txt")
	if s1 != 15 || s2 != 12 {
		t.Fatal("Wrong answer")
	}
}
