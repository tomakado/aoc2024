package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type report []int

func main() {
	reports := readInput()
	fmt.Println(numSafeReports(reports))
}

func numSafeReports(reports []report) int {
	var count int

	for _, r := range reports {
		if isSafeReport(r) {
			count++
			continue
		}

		for j := 0; j < len(r); j++ {
			cutReport := make([]int, 0, len(r)-1)

			for k := 0; k < len(r); k++ {
				if k == j {
					continue
				}
				cutReport = append(cutReport, r[k])
			}

			if isSafeReport(cutReport) {
				count++
				break
			}
		}
	}

	return count
}

func isSafeReport(r report) bool {
	return isAdjSafeReport(r) && (isDecreasing(r) || isIncreasing(r))
}

func isAdjSafeReport(r report) bool {
	if len(r) <= 1 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if !isAdjSafe(r[i-1], r[i]) {
			return false
		}
	}

	return true
}

func isAdjSafe(a, b int) bool {
	diff := abs(a - b)
	// fmt.Println("diff=", diff, "a=", a, "b=", b)
	return diff >= 1 && diff <= 3
}

func readInput() []report {
	var reports []report

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		reports = append(reports, readReport(line))
	}

	return reports
}

func readReport(line string) report {
	var levels report

	for _, v := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(v)
		levels = append(levels, n)
	}

	return levels
}

func isIncreasing(r report) bool {
	if len(r) <= 1 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if r[i] <= r[i-1] {
			return false
		}
	}

	return true
}

func isDecreasing(r report) bool {
	if len(r) <= 1 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if r[i] >= r[i-1] {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func removeByIndex(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}
