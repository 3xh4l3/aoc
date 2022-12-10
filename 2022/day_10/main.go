package main

import (
	"fmt"

	"github.com/3xh4l3/aoc"
)

func main() {
	p1, p2 := day_10("input.txt")
	fmt.Printf("\nPart one: %d", p1)
	fmt.Printf("\nPart two: %d\n", p2)
}

func day_10(input string) (p1, p2 int) {
	f, scanner := aoc.GetFile(input)
	defer f.Close()

	x := 1
	skip := false
	line := ""
	for i := 1; i <= 240; i++ {
		// Collect part one
		if i%40 == 20 {
			p1 += i * x
		}
		// Addx in process
		if !skip {
			scanner.Scan()
			line = scanner.Text()
		}
		// Addx operation process
		if line != "noop" {
			skip = !skip
			if !skip {
				var add int
				fmt.Sscanf(line, "addx %d", &add)
				x += add
			}
		}

		// Draw part two
		if i%40 == 1 && i != 1 {
			fmt.Printf(" i:%d\n", i)
		}
		if i%40 == x-1 || i%40 == x || i%40 == x+1 {
			fmt.Print("⬜")
		} else {
			fmt.Print("⬛")
		}
	}
	return
}
