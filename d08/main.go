package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var size int

func main() {
	freqs := initLocation(input)
	fmt.Println(countDistinctAntinodeLocations(freqs))
	fmt.Println(countDistinctAntinodeLocations2(freqs))
}

type vec2 [2]int

func (v vec2) x() int { return v[0] }
func (v vec2) y() int { return v[1] }

func (v vec2) turnRight() vec2 {
	return vec2{-v.y(), v.x()}
}

func (v vec2) add(other vec2) vec2 {
	return vec2{v.x() + other.x(), v.y() + other.y()}
}

func (v vec2) sub(other vec2) vec2 {
	return vec2{v.x() - other.x(), v.y() - other.y()}
}

func (v vec2) String() string {
	return fmt.Sprintf("%d:%d", v.y()+1, v.x()+1)
}

func countDistinctAntinodeLocations(freqs map[rune][]vec2) int {
	antinodes := make(map[vec2]struct{})

	for _, points := range freqs {
		for i := range points {
			for j := range points {
				if i == j {
					continue
				}

				diff := points[i].sub(points[j])
				opposite := points[i].add(diff)
				if !isInsideLocation(size, opposite) {
					continue
				}

				antinodes[opposite] = struct{}{}
			}
		}
	}

	return len(antinodes)
}

func countDistinctAntinodeLocations2(freqs map[rune][]vec2) int {
	antinodes := make(map[vec2]struct{})

	for _, points := range freqs {
		for i := range points {
			for j := range points {
				if i == j {
					continue
				}

				antinodes[points[i]] = struct{}{}
				antinodes[points[j]] = struct{}{}

				diagonal := buildDiagonal(points[i], points[j])
				for _, p := range diagonal {
					antinodes[p] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func buildDiagonal(a, b vec2) []vec2 {
	var d []vec2

	diff := a.sub(b)
	opposite := a.add(diff)

	for isInsideLocation(size, opposite) {
		d = append(d, opposite)
		opposite = opposite.add(diff)
	}

	return d
}

func isInsideLocation(size int, p vec2) bool {
	if p.y() < 0 || p.y() > size-1 {
		return false
	}

	return p.x() >= 0 && p.x() <= size-1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func initLocation(input string) map[rune][]vec2 {
	lines := strings.Split(input, "\n")
	freqs := make(map[rune][]vec2)

	for y, line := range lines {
		if line == "" {
			continue
		}

		for x, r := range line {
			if r == '.' {
				continue
			}

			points := freqs[r]
			freqs[r] = append(points, vec2{x, y})
		}

		size++
	}

	return freqs
}
