package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	rules   map[int]map[int]struct{}
	updates [][]int
)

func main() {
	readInput()

	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	var sum int

	for _, upd := range updates {
		if isCorrectUpdate(upd) {
			sum += upd[len(upd)/2]
		}
	}

	return sum
}

func partTwo() int {
	var sum int

	for _, upd := range updates {
		if !isCorrectUpdate(upd) {
			sorted := sortUpdate(upd)
			sum += sorted[len(sorted)/2]
		}
	}

	return sum
}

func isCorrectUpdate(update []int) bool {
	tracked := make(map[int]struct{})

	for _, p := range update {
		rule := rules[p]

		for p := range rule {
			if _, ok := tracked[p]; ok {
				return false
			}
		}

		tracked[p] = struct{}{}
	}

	return true
}

func sortUpdate(update []int) []int {
	sorted := make([]int, len(update))
	copy(sorted, update)

	sort.SliceStable(sorted, func(i, j int) bool {
		ruleJ, ruleJExists := rules[sorted[j]]
		if !ruleJExists {
			return true
		}

		_, contains := ruleJ[sorted[i]]
		return contains
	})

	return sorted
}

func readInput() {
	rulesInput, updatesInput, _ := strings.Cut(input, "\n\n")
	rules, updates = readRulesInput(rulesInput), readUpdatesInput(strings.TrimSuffix(updatesInput, "\n"))
}

func readRulesInput(in string) map[int]map[int]struct{} {
	r := make(map[int]map[int]struct{}, len(in))

	for _, ruleInput := range strings.Split(in, "\n") {
		aStr, bStr, _ := strings.Cut(ruleInput, "|")
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)

		rule, ok := r[a]
		if !ok {
			rule = make(map[int]struct{})
		}
		rule[b] = struct{}{}
		r[a] = rule
	}

	return r
}

func readUpdatesInput(in string) [][]int {
	u := make([][]int, 0, len(in))

	for _, updateInput := range strings.Split(in, "\n") {
		updateStr := strings.Split(updateInput, ",")
		update := make([]int, 0, len(updateStr))

		for _, pageNumStr := range updateStr {
			pageNum, _ := strconv.Atoi(pageNumStr)
			update = append(update, pageNum)
		}

		u = append(u, update)
	}

	return u
}
