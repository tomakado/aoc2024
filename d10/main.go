package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	tmap := readInput(input)
	fmt.Println(countTrailScore(tmap))
	fmt.Println(countTrailRating(tmap))
}

type vec2 [2]int

func (v vec2) x() int { return v[0] }
func (v vec2) y() int { return v[1] }

func (v vec2) add(other vec2) vec2 {
	return vec2{v.x() + other.x(), v.y() + other.y()}
}

func (v vec2) String() string {
	return fmt.Sprintf("%d:%d", v.y()+1, v.x()+1)
}

func countTrailRating(tmap [][]int) int {
	var rating int

	for y, line := range tmap {
		for x, h := range line {
			if h != 0 {
				continue
			}

			rating += countAllTrailsFrom(vec2{x, y}, tmap, 0)
		}
	}
	return rating
}

func countTrailScore(tmap [][]int) int {
	var score int

	for y, line := range tmap {
		for x, h := range line {
			if h != 0 {
				continue
			}

			score += countUniqueTrailsFrom(vec2{x, y}, make(map[vec2]struct{}), tmap, 0)
		}
	}

	return score
}

func countAllTrailsFrom(p vec2, tmap [][]int, acc int) int {
	currentHeight := tmap[p.y()][p.x()]
	if currentHeight == 9 {
		return acc + 1
	}

	newAcc := acc

	for _, dir := range directions {
		lookupP := p.add(dir)
		if !isInsideLocation(len(tmap), lookupP) {
			continue
		}

		lookupHeight := tmap[lookupP.y()][lookupP.x()]
		diff := lookupHeight - currentHeight

		if diff != 1 {
			continue
		}

		newAcc = countAllTrailsFrom(lookupP, tmap, newAcc)
	}

	return newAcc
}

var (
	up    = vec2{0, -1}
	right = vec2{1, 0}
	down  = vec2{0, 1}
	left  = vec2{-1, 0}
)

var directions = []vec2{up, right, down, left}

func countUniqueTrailsFrom(p vec2, visited map[vec2]struct{}, tmap [][]int, acc int) int {
	currentHeight := tmap[p.y()][p.x()]
	if currentHeight == 9 {
		if _, ok := visited[p]; ok {
			return acc
		}

		visited[p] = struct{}{}
		return acc + 1
	}

	newAcc := acc

	for _, dir := range directions {
		lookupP := p.add(dir)
		if !isInsideLocation(len(tmap), lookupP) {
			continue
		}

		lookupHeight := tmap[lookupP.y()][lookupP.x()]
		diff := lookupHeight - currentHeight

		if diff != 1 {
			continue
		}

		newAcc = countUniqueTrailsFrom(lookupP, visited, tmap, newAcc)
	}

	return newAcc
}

func isInsideLocation(size int, p vec2) bool {
	if p.y() < 0 || p.y() > size-1 {
		return false
	}

	return p.x() >= 0 && p.x() <= size-1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func readInput(input string) [][]int {
	var (
		lines = strings.Split(input, "\n")
		tmap  = make([][]int, 0, len(lines))
	)

	for _, line := range lines {
		if line == "" {
			continue
		}
		heights := make([]int, 0, len(line))

		for _, h := range line {
			height, _ := strconv.Atoi(string(h))
			heights = append(heights, height)
		}

		tmap = append(tmap, heights)
	}

	return tmap
}
