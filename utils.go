package aoc

import (
	"bufio"
	"log"
	"os"
)

func GetFile(name string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()

	return file, bufio.NewScanner(file)
}
