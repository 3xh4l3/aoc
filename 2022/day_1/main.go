package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var kcals, all []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		kcal, err := strconv.Atoi(scanner.Text())

		// err = empty line
		if err != nil {
			all = append(all, sum_kcals(kcals))
			kcals = make([]int, 0)
		} else {
			kcals = append(kcals, kcal)
		}
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i] > all[j]
	})

	// Part One - Find the Elf carrying the most Calories.
	fmt.Printf("Max carrying Calories: \t\t\t%d\n", all[0])

	// Part Two - Find the top three Elves carrying the most Calories.
	fmt.Printf("Sum of max three carrying Calories: \t%d\n", sum_kcals(all[0:3]))
}

func sum_kcals(kcals []int) (sum int) {
	for _, n := range kcals {
		sum += n
	}
	return
}
