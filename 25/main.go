package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("d01p1:", d1(readLines("input/01.txt")))
	//fmt.Println("d01p2:", d2(readLines("input/01test.txt")))
	//fmt.Println("(5 - 10) % 100 =", ((5-1095)%100+100)%100)
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic("failed to open: " + path)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
