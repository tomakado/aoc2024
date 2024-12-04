package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	m := readInput()
	fmt.Println(countXMAS(m))
	fmt.Println(countCrossMAS(m))
}

var xmasDirections = []struct{ v, h int }{
	{1, 0},   // right
	{-1, 0},  // left
	{0, -1},  // up
	{0, 1},   // down
	{1, 1},   // right-down
	{1, -1},  // right-up
	{-1, 1},  // left-down
	{-1, -1}, // left-up
}

func countXMAS(m [][]rune) int {
	var count int

	for i, line := range m {
		for j := range line {
			for _, dir := range xmasDirections {
				if checkXMAS(m, i, j, dir.v, dir.h) {
					count++
				}
			}
		}
	}

	return count
}

func checkXMAS(m [][]rune, posv, posh, dirv, dirh int) bool {
	const word = "XMAS"

	for i, ch := range word {
		v := posv + i*dirv
		h := posh + i*dirh

		switch {
		case dirv > 0 && len(m)-posv < len(word),
			dirv < 0 && posv < len(word)-1,
			dirh > 0 && len(m[posh])-posh < len(word),
			dirh < 0 && posh < len(word)-1,
			m[v][h] != ch:
			return false
		}
	}

	return true
}

func countCrossMAS(m [][]rune) int {
	var count int

	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			if m[i][j] != 'A' {
				continue
			}

			d1 := string(m[i-1][j-1]) + string(m[i+1][j+1])
			d2 := string(m[i-1][j+1]) + string(m[i+1][j-1])

			if (d1 == "MS" || d1 == "SM") && (d2 == "MS" || d2 == "SM") {
				count++
			}
		}
	}

	return count
}

func readInput() [][]rune {
	var chars [][]rune

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		s := make([]rune, 0, len(line))

		for _, ch := range line {
			s = append(s, ch)
		}

		chars = append(chars, s)
	}

	return chars
}
