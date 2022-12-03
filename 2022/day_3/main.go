package main

import (
	"fmt"
	"github.com/3xh4l3/aoc"
)

func main() {
	sum1, sum2 := day_3("input.txt")
	fmt.Printf("Part one: %d\n", sum1)
	fmt.Printf("Part two: %d\n", sum2)
}

func day_3(input string) (sum1, sum2 int) {
	f, scanner := aoc.GetFile(input)
	defer f.Close()

	var i int
	// Elves groups
	ss := make([]string, 3)
	for scanner.Scan() {
		line := scanner.Text()
		s1 := line[0 : len(line)/2]
		s2 := line[len(line)/2 : len(line)]

		// Part one
		sum1 += get_points(get_common_char([]string{s1, s2}))

		// Part two
		ss[i] = line
		i++
		if i > 2 {
			i = 0
			sum2 += get_points(get_common_char(ss))
		}

	}
	return
}

/*
Compare with hash logic
*/
func get_common_char(ss []string) (c rune) {
	// Make list of hashes
	lm := make([]map[rune]bool, len(ss))
	// Init hashes in list insted nil maps
	for i, _ := range lm {
		lm[i] = make(map[rune]bool)
	}
	// Ge first common char in strings
	for i, s := range ss {
		if i == 0 {
			for _, c := range s {
				lm[0][c] = true
			}
			continue
		}
		for _, c := range s {
			if lm[i-1][c] {
				if i == len(ss)-1 {
					return c
				}
				lm[i][c] = true
			}
		}
	}
	return
}

/*
ASCII A = 65
ASCII a = 97
*/
func get_points(i rune) (p int) {
	if i < 97 {
		p = int(i) - 38
	} else {
		p = int(i) - 96
	}
	return
}
