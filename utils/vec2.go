package utils

import "fmt"

type Vec2 struct {
	X, Y int
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) String() string {
	return fmt.Sprintf("%d:%d", v.Y+1, v.X+1)
}

func (v Vec2) IsInside(locsize int) bool {
	if v.Y < 0 || v.Y > locsize-1 {
		return false
	}

	return v.X >= 0 && v.X <= locsize-1
}

var (
	Up    = Vec2{0, -1}
	Right = Vec2{1, 0}
	Down  = Vec2{0, 1}
	Left  = Vec2{-1, 0}
)

var Directions = []Vec2{Up, Right, Down, Left}
