package main

import (
	"strconv"
)

func d1(lines []string) string {
	cur, res := 50, 0
	for _, rotation := range lines {
		cur = rotate(cur, rotation)
		if cur == 0 {
			res++
		}
	}
	return strconv.Itoa(res)
}

func d2(lines []string) string {
	if len(lines) > 0 {
	}
	return "answer"
}

func rotate(cur int, move string) int {
	n, _ := strconv.Atoi(move[1:])
	if move[0] == 'R' {
		cur += n
	} else {
		cur -= n
	}
	return (cur%100 + 100) % 100
}
