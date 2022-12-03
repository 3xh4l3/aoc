package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	s1, s2 := day_2("input.txt")
	fmt.Printf("Total score for Part 1: %d\n", s1)
	fmt.Printf("Total score for Part 2: %d\n", s2)
}

func day_2(filename string) (int, int) {
	file, err := os.Open(filename)
	//file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Points
	t1 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	t2 := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	s := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	var s1, s2 int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, " ")
		a := t1[pair[0]]
		b := t2[pair[1]]
		r := s[pair[1]]
		// Part 1 - Get result by shape
		s1 += b + get_result(a, b)
		// Part 2 - Guess shape by result
		s2 += r + get_shape(a, r)
	}
	return s1, s2
}

func get_result(a, b int) (score int) {
	switch {
	case (a == 1 && b == 1) || (a == 2 && b == 2) || (a == 3 && b == 3):
		score = 3
	case (a == 1 && b == 2) || (a == 2 && b == 3) || (a == 3 && b == 1):
		score = 6
	}
	return
}

func get_shape(a, b int) (shape int) {
	shape = 1
	switch {
	case (a == 1 && b == 6) || (a == 2 && b == 3) || (a == 3 && b == 0):
		shape = 2
	case (a == 1 && b == 0) || (a == 2 && b == 6) || (a == 3 && b == 3):
		shape = 3
	}
	return
}
