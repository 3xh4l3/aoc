package main

import (
	"fmt"

	"github.com/3xh4l3/aoc"
)

func main() {
	p1, p2 := day_8("input.txt")
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}

func day_8(input string) (p1, p2 int) {
	var matrix []string

	f, scanner := aoc.GetFile(input)
	defer f.Close()

	// Get matrix
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	for i, r := range matrix {
		for j, c := range r {
			// First - visibility distanse
			// Second - tree visibility from outside
			fu, up := lookUp(matrix, i, j, c)
			fd, down := lookDown(matrix, i, j, c)
			fl, left := lookLeft(r, j, c)
			fr, right := lookRight(r, j, c)

			// Part one
			if up || down || left || right {
				p1++
			}

			// Part two
			m := fu * fd * fl * fr
			if m > p2 {
				p2 = m
			}
		}
	}

	return
}

func lookUp(m []string, h, v int, c rune) (int, bool) {
	i := 1
	for ; i <= h; i++ {
		if rune(m[h-i][v]) >= c {
			return i, false
		}
	}
	return i - 1, true
}

func lookDown(m []string, h, v int, c rune) (int, bool) {
	i := 1
	for ; i < len(m)-h; i++ {
		if rune(m[h+i][v]) >= c {
			return i, false
		}
	}
	return i - 1, true
}

func lookLeft(s string, j int, c rune) (int, bool) {
	i := 1
	for ; i <= j; i++ {
		if rune(s[j-i]) >= c {
			return i, false
		}
	}
	return i - 1, true
}

func lookRight(s string, j int, c rune) (int, bool) {
	i := 1
	for ; i < len(s)-j; i++ {
		if rune(s[j+i]) >= c {
			return i, false
		}
	}
	return i - 1, true
}
