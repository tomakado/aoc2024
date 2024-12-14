package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	var (
		diskmap     = readInput(input)
		unwrapped   = unwrap(diskmap)
		compressed  = compress(unwrapped)
		compressed2 = compress2(unwrapped)
	)

	fmt.Println(checksum(compressed))
	fmt.Println(checksum(compressed2))
}

var (
	fileSizes     = make(map[int]int)
	fileLocations = make(map[int]int)
)

func compress2(unwrapped []int) []int {
	var (
		compressed = make([]int, len(unwrapped))
		fids       = fileIDsSorted()
		maxFileID  = fids[0]
	)

	copy(compressed, unwrapped)

	for id := maxFileID; id >= 0; id-- {
		var freeSpace, cursor int

		for cursor < fileLocations[id] && freeSpace < fileSizes[id] {
			cursor += freeSpace
			freeSpace = 0

			for compressed[cursor] != -1 {
				cursor++
			}

			for cursor+freeSpace < len(compressed) && compressed[cursor+freeSpace] == -1 {
				freeSpace++
			}
		}

		if cursor >= fileLocations[id] {
			continue
		}

		for i := cursor; i < cursor+fileSizes[id]; i++ {
			compressed[i] = id
		}

		for i := fileLocations[id]; i < fileLocations[id]+fileSizes[id]; i++ {
			compressed[i] = -1
		}
	}

	return compressed
}

func fileIDsSorted() []int {
	fids := make([]int, 0, len(fileSizes))
	for k := range fileSizes {
		fids = append(fids, k)
	}

	sort.Slice(fids, func(i, j int) bool {
		return fids[i] > fids[j]
	})

	return fids
}

func compress(unwrapped []int) []int {
	compressed := make([]int, len(unwrapped))
	copy(compressed, unwrapped)

	var cursor int
	for compressed[cursor] != -1 {
		cursor++
	}

	backCursor := len(compressed) - 1
	for compressed[backCursor] == -1 {
		backCursor--
	}

	for backCursor > cursor {
		compressed[cursor] = compressed[backCursor]
		compressed[backCursor] = -1

		for compressed[backCursor] == -1 {
			backCursor--
		}

		for compressed[cursor] != -1 {
			cursor++
		}
	}

	return compressed
}

func checksum(compressed []int) int {
	var sum int

	for i, d := range compressed {
		if d == -1 {
			continue
		}
		sum += i * d
	}

	return sum
}

func printUnwrapped(unwrapped []int) {
	for _, u := range unwrapped {
		if u == -1 {
			fmt.Print(".")
			continue
		}

		fmt.Print(u)
	}

	fmt.Println()
}

func printCompressed(compressed []int) {
	for _, d := range compressed {
		if d == -1 {
			fmt.Print(".")
			continue
		}

		fmt.Print(d)
	}

	fmt.Println()
}

func unwrap(diskmap []int) []int {
	var (
		unwrapped []int
		id        int
	)

	for i, d := range diskmap {
		if i%2 == 0 {
			fileLocations[id] = len(unwrapped)
			fileSizes[id] = d
			for j := 0; j < d; j++ {
				unwrapped = append(unwrapped, id)
			}
			id++
			continue
		}

		for j := 0; j < d; j++ {
			unwrapped = append(unwrapped, -1)
		}
	}

	return unwrapped
}

func readInput(input string) []int {
	digits := make([]int, 0, len(input))

	for _, r := range input {
		if r == '\n' {
			continue
		}

		digit, _ := strconv.Atoi(string(r))
		digits = append(digits, digit)
	}

	return digits
}
