package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/3xh4l3/aoc"
)

const DELAY = 50

func main() {
	p1, p2 := day_9("input.txt")
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}

func day_9(input string) (p1, p2 int) {
	f, scanner := aoc.GetFile(input)
	defer f.Close()

	// Init
	var xh, yh int

	// Knots
	knots := make([][2]int, 9)
	// Store uniq coords of some knot
	uniq_coords_1 := make(map[string]bool)
	uniq_coords_9 := make(map[string]bool)

	// Map fo actions
	m := map[string]func(int, int) (int, int){
		"R": func(x, y int) (int, int) { return x + 1, y },
		"L": func(x, y int) (int, int) { return x - 1, y },
		"U": func(x, y int) (int, int) { return x, y + 1 },
		"D": func(x, y int) (int, int) { return x, y - 1 },
	}

	for scanner.Scan() {
		var (
			direction string
			count     int
		)
		line := scanner.Text()
		fmt.Sscanf(line, "%s %d", &direction, &count)

		for i := 1; i <= count; i++ {
			// Change Head position
			xh, yh = m[direction](xh, yh)
			// Update hole rope
			for i, knot := range knots {
				// Reasign head coords to previous knot
				var x, y int
				if i == 0 {
					// Get head on first step
					x, y = xh, yh
				} else {
					// Get previous knot
					x, y = knots[i-1][0], knots[i-1][1]
				}
				knots[i][0], knots[i][1] = changeKnotCoords(x, y, knot[0], knot[1])
			}
			drawGrid(xh, yh, knots)
			uniq_coords_1[strconv.Itoa(knots[0][0])+strconv.Itoa(knots[0][1])] = true
			uniq_coords_9[strconv.Itoa(knots[8][0])+strconv.Itoa(knots[8][1])] = true
		}
	}
	p1 = len(uniq_coords_1)
	p2 = len(uniq_coords_9)

	return
}

// Visualize
func drawGrid(xh, yh int, knots [][2]int) {
	scale := 30
	time.Sleep(DELAY * time.Millisecond)
	for i := yh + scale; i >= yh-scale; i -= 1 {
		for j := xh - scale; j <= xh+scale; j++ {
			s := ". "
			if xh == j && yh == i {
				s = "H "
			}
			for k, knot := range knots {
				if knot[0] == j && knot[1] == i {
					s = fmt.Sprintf("%d ", k)
				}
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

// Motions
func changeKnotCoords(xh, yh, xt, yt int) (int, int) {
	switch {
	// Diag upright
	case (xh > xt && yh-yt > 1) || (xh-xt > 1 && yh > yt):
		return xt + 1, yt + 1
	// Diag upleft
	case (xh < xt && yh-yt > 1) || (xt-xh > 1 && yh > yt):
		return xt - 1, yt + 1
	// Diag downright
	case (xh > xt && yt-yh > 1) || (xh-xt > 1 && yh < yt):
		return xt + 1, yt - 1
	// Diag downleft
	case (xh < xt && yt-yh > 1) || (xt-xh > 1 && yh < yt):
		return xt - 1, yt - 1

	// Right
	case xh-xt > 1 && yh == yt:
		return xt + 1, yt
	// Up
	case xh == xt && yh-yt > 1:
		return xt, yt + 1
	// Left
	case xt-xh > 1 && yt == yh:
		return xt - 1, yt
	// Down
	case xt == xh && yt-yh > 1:
		return xt, yt - 1
	default:
		return xt, yt
	}
}
