package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/tomakado/aoc2024/utils"
)

//go:embed input.txt
var input string

func main() {
	var (
		garden  = readInput(input)
		regions = collectRegions(garden)
	)

	fmt.Println(totalFencingPrice(regions))
	fmt.Println(totalFencingPrice2(regions))
}

type region struct {
	plant rune
	plots []utils.Vec2
}

func (r region) fencingPrice() int {
	return len(r.plots) * r.perimeter()
}

func (r region) fencingPrice2() int {
	return len(r.plots) * r.numOfSides()
}

func (r region) perimeter() int {
	var (
		perimeter int
		pointsMap = r.pointsMap()
	)

	for _, p := range r.plots {
		for _, dir := range utils.Directions {
			lookupPos := p.Add(dir)

			if _, ok := pointsMap[lookupPos]; !ok {
				perimeter++
			}
		}
	}

	return perimeter
}

func (r region) numOfSides() int {
	var (
		pointsMap = r.pointsMap()
		checks    = [][]utils.Vec2{
			{utils.Left, utils.Up},
			{utils.Down, utils.Left},
			{utils.Right, utils.Down},
			{utils.Up, utils.Right},
		}
		numSides int
	)

	for _, p := range r.plots {
		for _, check := range checks {
			dirA, dirB := check[0], check[1]
			pointA := p.Add(dirA)
			pointB := p.Add(dirB)
			pointAB := p.Add(dirA).Add(dirB)
			_, aOk := pointsMap[pointA]
			_, bOk := pointsMap[pointB]
			_, abOk := pointsMap[pointAB]

			if (!aOk && !bOk) || (aOk && bOk && !abOk) {
				numSides++
			}
		}
	}

	return numSides
}

func (r region) pointsMap() map[utils.Vec2]struct{} {
	pointsMap := make(map[utils.Vec2]struct{}, len(r.plots))
	for _, p := range r.plots {
		pointsMap[p] = struct{}{}
	}

	return pointsMap
}

func (r region) String() string {
	var out strings.Builder

	out.WriteRune(r.plant)
	out.WriteString("; ( ")

	for _, plot := range r.plots {
		out.WriteString(plot.String())
		out.WriteString(" ")
	}

	out.WriteString("); perimeter = ")
	out.WriteString(strconv.Itoa(r.perimeter()))

	return out.String()
}

func totalFencingPrice(regions []region) int {
	var totalPrice int

	for _, reg := range regions {
		totalPrice += reg.fencingPrice()
	}

	return totalPrice
}

func totalFencingPrice2(regions []region) int {
	var totalPrice int

	for _, reg := range regions {
		totalPrice += reg.fencingPrice2()
	}

	return totalPrice
}

func collectRegions(garden [][]rune) []region {
	var regions []region

	for y, line := range garden {
		for x, plant := range line {
			var (
				pos = utils.Vec2{X: x, Y: y}
				acc = region{plant: plant}
			)

			newReg := collectRegion(garden, acc, pos, utils.Vec2{})
			if len(newReg.plots) == 0 {
				continue
			}

			regions = append(regions, newReg)
		}
	}

	return regions
}

var visited = make(map[utils.Vec2]struct{})

func collectRegion(garden [][]rune, acc region, pos, incomingDirection utils.Vec2) region {
	if _, ok := visited[pos]; ok {
		return acc
	}

	visited[pos] = struct{}{}
	newAcc := acc
	newAcc.plots = append(newAcc.plots, pos)

	for _, dir := range utils.Directions {
		lookupPos := pos.Add(dir)

		if lookupPos == incomingDirection {
			continue
		}

		if lookupPos.IsInside(len(garden)) && garden[lookupPos.Y][lookupPos.X] == newAcc.plant {
			newAcc = collectRegion(
				garden,
				newAcc,
				lookupPos,
				pos,
			)
			continue
		}
	}

	return newAcc
}

func readInput(input string) [][]rune {
	var garden [][]rune

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		garden = append(garden, []rune(line))
	}

	return garden
}
