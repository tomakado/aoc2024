package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type equation struct {
	left  int
	right []int
}

func main() {
	eqs := readInput()
	fmt.Println(sumLeftPossiblyTrue(eqs))
}

func sumLeftPossiblyTrue(eqs []equation) int {
	var sum int

	for _, eq := range eqs {
		if isPossiblyTrue(eq.left, eq.right, 0) {
			sum += eq.left
		}
	}

	return sum
}

func isPossiblyTrue(left int, right []int, acc int) bool {
	if len(right) == 1 {
		return acc+right[0] == left || acc*right[0] == left || concat(acc, right[0]) == left
	}

	return isPossiblyTrue(left, right[1:], acc+right[0]) ||
		isPossiblyTrue(left, right[1:], acc*right[0]) ||
		isPossiblyTrue(left, right[1:], concat(acc, right[0]))
}

func concat(a, b int) int {
	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return res
}

func readInput() []equation {
	lines := strings.Split(input, "\n")
	eqs := make([]equation, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		leftStr, rightStr, _ := strings.Cut(line, ": ")
		left, _ := strconv.Atoi(leftStr)

		rightElems := strings.Split(rightStr, " ")
		right := make([]int, 0, len(rightElems))

		for _, e := range rightElems {
			intElem, _ := strconv.Atoi(e)
			right = append(right, intElem)
		}

		eqs = append(eqs, equation{left, right})
	}

	return eqs
}
