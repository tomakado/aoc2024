package main

import (
	"fmt"
	"strconv"
	"strings"

	"go.tomakado.io/containers/list"
)

const input = "70949 6183 4 3825336 613971 0 15 182"

func main() {
	l := deserialize(input)

	for i := 0; i < 25; i++ {
		blink(l)
	}

	fmt.Println(l.Len())
}

func blink(l *list.List[int]) {
	for e := l.Front(); e != nil; e = e.Next() {
		switch {
		case e.Value == 0:
			e.Value = 1
		case numDigits(e.Value)%2 == 0:
			left, right := splitNum(e.Value)
			e.Value = left
			l.InsertAfter(right, e)
			e = e.Next()
		default:
			e.Value *= 2024
		}
	}
}

func serialize(l *list.List[int]) string {
	var out strings.Builder

	for e := l.Front(); e != nil; e = e.Next() {
		out.WriteString(strconv.Itoa(e.Value))
		out.WriteString(" ")
	}

	return out.String()
}

func numDigits(num int) int {
	var numDigits int

	x := num
	for x > 0 {
		numDigits++
		x /= 10
	}

	return numDigits
}

func splitNum(num int) (int, int) {
	var (
		n = numDigits(num)
		p = intPow(10, n/2)
	)
	return num / p, num % p
}

func deserialize(input string) *list.List[int] {
	l := list.New[int]().Init()

	for _, e := range strings.Split(input, " ") {
		if e == "" {
			continue
		}

		num, _ := strconv.Atoi(e)
		l.PushBack(num)
	}

	return l
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
