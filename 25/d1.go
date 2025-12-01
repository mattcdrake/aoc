package main

import (
	"strconv"
)

func d1(lines []string) string {
	cur, res := 50, 0
	for _, rotation := range lines {
		cur, _ = rotate(cur, rotation)
		if cur == 0 {
			res++
		}
	}
	return strconv.Itoa(res)
}

func d2(lines []string) string {
	cur, res, crosses := 50, 0, 0
	for _, rotation := range lines {
		cur, crosses = rotate(cur, rotation)
		res += crosses
	}
	return strconv.Itoa(res)
}

func rotate(cur int, move string) (int, int) {
	dist, _ := strconv.Atoi(move[1:])
	init := cur
	zeroes := dist / 100
	dist %= 100

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
