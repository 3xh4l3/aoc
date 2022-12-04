package main

import (
	"fmt"
	"github.com/3xh4l3/aoc"
	"log"
	"strconv"
	"strings"
)

func main() {
	p1, p2 := day_4("input.txt")
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}

func day_4(input string) (p1, p2 int) {
	f, scanner := aoc.GetFile(input)
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()

		pair := strings.Split(line, ",")
		r1 := fmap(strings.Split(pair[0], "-"))
		r2 := fmap(strings.Split(pair[1], "-"))

		// Part one
		if (r1[0] <= r2[0] && r1[1] >= r2[1]) || (r1[0] >= r2[0] && r1[1] <= r2[1]) {
			p1++
		}

		// Part two
		if (r1[0] <= r2[0] && r1[1] >= r2[0]) || (r1[0] <= r2[1] && r1[1] >= r2[1]) || (r1[0] >= r2[0] && r1[1] <= r2[1]) {
			p2++
		}
	}

	return
}

// Converts slice of string to slice of int
func fmap(s []string) (i []int) {
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		i = append(i, n)
	}
	return
}
