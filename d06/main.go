package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	initRoom()
	fmt.Println(countDistinctPositions())

	initRoom()
	fmt.Println(countLoopers())
}

type trail struct {
	rot, dst vec2
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

func (v vec2) String() string {
	return fmt.Sprintf("(%d, %d)", v.x(), v.y())
}

var (
	room          [][]rune
	startPos      vec2
	currentPos    vec2
	currentRot    = up
	visited       = make(map[vec2]struct{})
	visitedTrails = make(map[trail]struct{})
)

var (
	up    = vec2{0, -1}
	right = vec2{1, 0}
	down  = vec2{0, 1}
	left  = vec2{-1, 0}
)

func countDistinctPositions() int {
	nextTile := currentPos.add(currentRot)
	for isInsideRoom(nextTile) {
		move()
		nextTile = currentPos
	}

	return len(visited)
}

func countLoopers() int {

	var loopersCnt int

	roomCopy := make([][]rune, len(room))
	copy(roomCopy, room)

	for y, line := range roomCopy {
		for x, tile := range line {
			tilePos := vec2{x, y}
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("recovered at %s\n", tilePos)
				}
			}()

			if tilePos == startPos {
				continue
			}

			if tile == '#' {
				continue
			}

			initRoom()
			room[y] = putObstacleAt(room[y], x)
			if checkForLoop() {
				loopersCnt++
			}
		}
	}

	return loopersCnt
}

func checkForLoop() bool {
	for isInsideRoom(currentPos) {
		tr := trail{
			rot: currentRot,
			dst: currentPos.add(currentRot),
		}
		if _, ok := visitedTrails[tr]; ok {
			return true
		}

		if !move() {
			break
		}
	}

	return false
}

func isInsideRoom(v vec2) bool {
	if v.y() < 0 || v.y() > len(room)-1 {
		return false
	}

	return v.x() >= 0 && v.x() <= len(room[v.y()])-1
}

func move() bool {
	nextTile := currentPos.add(currentRot)

	// fmt.Println(currentPos, "->", nextTile)
	if !isInsideRoom(nextTile) {
		// fmt.Println(nextTile, "is outside the room")
		return false
	}

	for isObstacle(nextTile) {
		currentRot = currentRot.turnRight()
		nextTile = currentPos.add(currentRot)
	}

	currentPos = nextTile
	visited[currentPos] = struct{}{}

	tr := trail{rot: currentRot, dst: nextTile}
	visitedTrails[tr] = struct{}{}

	return true
}

func isObstacle(dst vec2) bool {
	return room[dst.y()][dst.x()] == '#'
}

func initRoom() {
	room = nil

	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		room = append(room, []rune(line))

		for x, tile := range line {
			if tile == '^' {
				startPos = vec2{x, y}
				currentPos = startPos
				currentRot = up

				visited = make(map[vec2]struct{})
				visited[currentPos] = struct{}{}

				visitedTrails = make(map[trail]struct{})
			}
		}
	}
}

func putObstacleAt(line []rune, pos int) []rune {
	var out strings.Builder

	for i, r := range line {
		if i == pos {
			out.WriteRune('#')
			continue
		}

		out.WriteRune(r)
	}

	return []rune(out.String())
}
