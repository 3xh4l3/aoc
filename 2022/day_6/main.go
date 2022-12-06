package main

import (
	"fmt"
	"github.com/3xh4l3/aoc"
	"log"
)

func main() {
	p1, p2 := day_6("input.txt")
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}

func day_6(input string) (p1, p2 int) {
	f, scanner := aoc.GetFile(input)
	defer f.Close()

	scanner.Scan()
	line := scanner.Text()

	// Part one
	p1, err := getUniqWindow(line, 4)
	if err != nil {
		log.Fatal(err)
	}

	// Part two
	p2, err = getUniqWindow(line, 14)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func getUniqWindow(s string, size int) (int, error) {
	for i := 0; i < (len(s) - size); i++ {
		if isuniq(s[i : i+size]) {
			return i + size, nil
		}
	}
	return 0, fmt.Errorf("Uniq window with size %d not found", size)
}

// O(n) uniq cheak by hash
func isuniq(s string) bool {
	m := make(map[rune]bool)
	for _, v := range s {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}
