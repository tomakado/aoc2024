package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tomakado/aoc2024/utils"
)

//go:embed input.txt
var input string

type App struct {
	frames           []map[utils.Vec2]int
	robots           []robotConfig
	currentRoomFrame int
}

func (a *App) init(robots []robotConfig, preSim int) {
	a.robots = robots
	a.frames = append(a.frames, robotsToFrame(robots))

	for i := 0; i < preSim; i++ {
		a.simulateAndSave()
	}
}

func robotsToFrame(robots []robotConfig) map[utils.Vec2]int {
	frame := make(map[utils.Vec2]int)

	for _, robot := range robots {
		frame[robot.pos]++
	}

	return frame
}

func (a *App) Update() error {
	switch {
	case repeatingKeyPressed(ebiten.KeySpace),
		repeatingKeyPressed(ebiten.KeyRight):

		if a.currentRoomFrame < len(a.frames)-1 {
			a.currentRoomFrame++
			return nil
		}

		a.simulateAndSave()

		return nil
	case repeatingKeyPressed(ebiten.KeyLeft):
		if a.currentRoomFrame > 0 {
			a.currentRoomFrame--
		}
		return nil
	default:
		return nil
	}
}

func (a *App) simulateAndSave() {
	a.robots = simulateOneSecond(a.robots)
	a.frames = append(a.frames, robotsToFrame(a.robots))
	a.currentRoomFrame++
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

const (
	tileSizeW = W * 0.01
	tileSizeH = H * 0.01
)

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	if len(a.frames) == 0 {
		return
	}

	room := a.frames[a.currentRoomFrame]

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if _, ok := room[utils.Vec2{X: x, Y: y}]; ok {
				vector.DrawFilledRect(
					screen,
					float32(x)*tileSizeW,
					float32(y)*tileSizeH,
					tileSizeW,
					tileSizeH,
					color.White,
					false,
				)
				continue
			}
		}
	}

	ebitenutil.DebugPrint(screen, strconv.Itoa(a.currentRoomFrame))
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return W, H
}

const (
	W = 101
	H = 103
)

func main() {
	sf, _ := safetyFactorAfter(readInput(input), 100)
	fmt.Println(sf)

	var preSim int
	args := os.Args
	if len(args) == 2 {
		var err error
		preSim, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	robots := readInput(input)

	ebiten.SetWindowSize(W*4, H*4)
	ebiten.SetWindowTitle("day 14")

	app := &App{}
	app.init(robots, preSim)

	if err := ebiten.RunGame(app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type robotConfig struct {
	pos, velocity utils.Vec2
}

func (c robotConfig) String() string {
	return fmt.Sprintf(
		"p=%d,%d v=%d,%d",
		c.pos.X,
		c.pos.Y,
		c.velocity.X,
		c.velocity.Y,
	)
}

func safetyFactorAfter(robots []robotConfig, seconds int) (int, []robotConfig) {
	robotsCopy := make([]robotConfig, len(robots))
	copy(robotsCopy, robots)

	for i := 0; i < seconds; i++ {
		robotsCopy = simulateOneSecond(robotsCopy)
	}

	// quadrants counters
	var first, second, third, fourth int

	for _, robot := range robotsCopy {
		switch {
		case robot.pos.X < W/2 && robot.pos.Y < H/2:
			first++
		case robot.pos.X > W/2 && robot.pos.Y < H/2:
			second++
		case robot.pos.X > W/2 && robot.pos.Y > H/2:
			third++
		case robot.pos.X < W/2 && robot.pos.Y > H/2:
			fourth++
		}
	}
	return first * second * third * fourth, robotsCopy
}

func simulateOneSecond(robots []robotConfig) []robotConfig {
	newRobots := make([]robotConfig, 0, len(robots))

	for _, robot := range robots {
		newPos := robot.pos.Add(robot.velocity)

		for newPos.X < 0 {
			newPos.X += W
		}

		for newPos.X >= W {
			newPos.X -= W
		}

		for newPos.Y < 0 {
			newPos.Y += H
		}

		for newPos.Y >= H {
			newPos.Y -= H
		}

		robot.pos = newPos
		newRobots = append(newRobots, robot)
	}

	return newRobots
}

func readInput(input string) []robotConfig {
	var configs []robotConfig

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var config robotConfig
		fmt.Sscanf(
			line,
			"p=%d,%d v=%d,%d",
			&config.pos.X,
			&config.pos.Y,
			&config.velocity.X,
			&config.velocity.Y,
		)

		configs = append(configs, config)
	}

	return configs
}
