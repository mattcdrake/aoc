package main

import (
	"strconv"
)

func d01p1(lines []string) int {
	cur, res := 50, 0
	for _, rotation := range lines {
		cur, _ = rotate(cur, rotation)
		if cur == 0 {
			res++
		}
	}
	return res
}

func d01p2(lines []string) int {
	cur, res, crosses := 50, 0, 0
	for _, rotation := range lines {
		cur, crosses = rotate(cur, rotation)
		res += crosses
	}
	return res
}

func rotate(init int, move string) (int, int) {
	dist, _ := strconv.Atoi(move[1:])
	zeroes := dist / 100
	dist %= 100

	cur := init
	if move[0] == 'R' {
		cur += dist
	} else {
		cur -= dist
	}

	if init != 0 && (cur <= 0 || cur > 99) {
		zeroes++
	}

	cur = (cur%100 + 100) % 100
	return cur, zeroes
}
