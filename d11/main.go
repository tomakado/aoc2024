package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "70949 6183 4 3825336 613971 0 15 182"

func main() {
	nums := readInput(input)
	fmt.Println(s(nums, 25))
	fmt.Println(s(nums, 75))
}

type memoTuple struct{ num, acc int }

var sMemo = map[memoTuple]int{}

func s(nums []int, acc int) int {
	if acc == 0 {
		return len(nums)
	}

	if len(nums) == 1 {
		num := nums[0]
		if val, ok := sMemo[memoTuple{num, acc}]; ok {
			return val
		}

		var toMemoize int

		switch {
		case num == 0:
			toMemoize = s([]int{1}, acc-1)
		case numDigits(num)%2 == 0:
			left, right := splitNum(num)
			toMemoize = s([]int{left, right}, acc-1)
		default:
			toMemoize = s([]int{num * 2024}, acc-1)
		}

		sMemo[memoTuple{num, acc}] = toMemoize
		return toMemoize
	}

	var sum int

	for _, n := range nums {
		sum += s([]int{n}, acc)
	}

	return sum
}

var numDigitsMemo = map[int]int{}

func numDigits(num int) int {
	if val, ok := numDigitsMemo[num]; ok {
		return val
	}

	var numDigits int

	x := num
	for x > 0 {
		numDigits++
		x /= 10
	}

	numDigitsMemo[num] = numDigits

	return numDigits
}

func splitNum(num int) (int, int) {
	var (
		n = numDigits(num)
		p = intPow(10, n/2)
	)
	return num / p, num % p
}

func readInput(input string) []int {
	var nums []int

	for _, e := range strings.Split(input, " ") {
		if e == "" {
			continue
		}

		num, _ := strconv.Atoi(e)
		nums = append(nums, num)
	}

	return nums
}

func intPow(n, m int) int {
	if n == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	p := n
	for i := 2; i <= m; i++ {
		p *= n
	}

	return p
}
