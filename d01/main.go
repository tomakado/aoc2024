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
	var (
		first, second = readInput()
	)

	fmt.Println(totalDistance(first, second))
	fmt.Println(similarityScore(first, second))
}

func totalDistance(f, s []int) int {
	var (
		first, second []int
		totalDistance int
	)

	first = append(first, f...)
	second = append(second, s...)

	for len(first) > 0 && len(second) > 0 {
		minX, minXIdx := minWithIdx(first)
		minY, minYIdx := minWithIdx(second)
		firstMinIndex, secondMinIndex := minXIdx, minYIdx

		d := abs(minX - minY)
		totalDistance += d

		first = removeByIndex(first, firstMinIndex)
		second = removeByIndex(second, secondMinIndex)
	}

	return totalDistance
}

func similarityScore(f, s []int) int {
	var totalScore int

	for _, v := range f {
		count := 0
		for _, w := range s {
			if v == w {
				count++
			}
		}

		totalScore += v * count
	}

	return totalScore
}

func readInput() ([]int, []int) {
	var first, second []int

	for _, line := range strings.Split(input, "\n") {
		fStr, sStr, _ := strings.Cut(line, "   ")

		f, _ := strconv.Atoi(fStr)
		s, _ := strconv.Atoi(sStr)

		first = append(first, f)
		second = append(second, s)
	}

	return first, second
}

func minWithIdx(arr []int) (int, int) {
	minIdx, min := 0, arr[0]

	for i, v := range arr {
		if v < min {
			minIdx, min = i, v
		}
	}

	return min, minIdx
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeByIndex(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}
