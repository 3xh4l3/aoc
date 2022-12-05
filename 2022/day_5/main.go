package main

import (
	"fmt"
	"github.com/3xh4l3/aoc"
	"strconv"
	"strings"
	"unicode"
)

func insert(stack []string, ins string) []string {
	return append([]string{ins}, stack...)
}

func pop(stack []string) ([]string, string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func push(stack []string, s string) []string {
	return append(stack, s)
}

func pop_stack(stack []string, c int) ([]string, []string) {
	return stack[:len(stack)-c], stack[len(stack)-c:]
}

func push_stack(stack, s []string) []string {
	return append(stack, s...)
}

func main() {
	p1, p2 := day_5("input.txt")
	fmt.Printf("Part one: %s\n", p1)
	fmt.Printf("Part two: %s\n", p2)
}

func day_5(input string) (p1, p2 string) {
	f, scanner := aoc.GetFile(input)
	defer f.Close()

	stacks_p1 := make([][]string, 0)
	stacks_p2 := make([][]string, 0)

	var parse_stacks bool = true
	var parse_actions bool

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, " 1") {
			parse_stacks = false
		}
		if strings.HasPrefix(line, "move") {
			parse_actions = true
		}

		// Get stacks
		if parse_stacks {
			// Count and init stacks
			if len(stacks_p1) < (len(line)+1)/4 {
				for i := 0; i <= len(line)/4; i++ {
					stacks_p1 = append(stacks_p1, make([]string, 0))
					stacks_p2 = append(stacks_p2, make([]string, 0))
				}
			}

			// Fill stacks
			for i, j := 0, 1; j < len(line); i, j = i+1, j+4 {
				if unicode.IsLetter(rune(line[j])) {
					stacks_p1[i] = insert(stacks_p1[i], string(line[j]))
					stacks_p2[i] = insert(stacks_p2[i], string(line[j]))
				}
			}
		}

		// Get actions
		if parse_actions {
			actions := strings.Split(line, " ")
			move, _ := strconv.Atoi(actions[1])
			from, _ := strconv.Atoi(actions[3])
			to, _ := strconv.Atoi(actions[5])

			from--
			to--

			// Make action
			// Part one
			var poped_1 string
			for i := 0; i < move; i++ {
				stacks_p1[from], poped_1 = pop(stacks_p1[from])
				stacks_p1[to] = push(stacks_p1[to], poped_1)
			}
			// Part two
			poped_2 := make([]string, move)
			stacks_p2[from], poped_2 = pop_stack(stacks_p2[from], move)
			stacks_p2[to] = push_stack(stacks_p2[to], poped_2)
		}
	}

	p_1 := make([]string, len(stacks_p1))
	p_2 := make([]string, len(stacks_p2))
	for _, v := range stacks_p1 {
		p_1 = append(p_1, v[len(v)-1])
	}
	for _, v := range stacks_p2 {
		p_2 = append(p_2, v[len(v)-1])
	}
	p1 = strings.Join(p_1, "")
	p2 = strings.Join(p_2, "")

	return
}
