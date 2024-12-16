package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	configs := readInput(input)

	fmt.Println(fewestTokensToWinAll(configs, 0))
	fmt.Println(fewestTokensToWinAll(configs, 10000000000000))
}

func fewestTokensToWinAll(configs []config, shift float64) float64 {
	var total float64

	for _, cfg := range configs {
		tokens, isWinnable := cfg.fewestTokensToWin(shift)
		if !isWinnable {
			continue
		}

		total += tokens
	}

	return total
}

type config struct {
	ButtonA, ButtonB, Prize floatvec2
}

type floatvec2 struct{ X, Y float64 }

func (cfg config) fewestTokensToWin(shift float64) (float64, bool) {
	cfg.Prize.X += shift
	cfg.Prize.Y += shift

	var (
		a = cfg.ButtonA.X
		b = cfg.ButtonB.X
		c = cfg.ButtonA.Y
		d = cfg.ButtonB.Y
		A = cfg.Prize.X
		B = cfg.Prize.Y

		det = a*d - b*c
	)
	if det == 0 {
		return 0, false
	}

	var (
		x = (A*d - B*b) / det
		y = (a*B - c*A) / det
	)

	if x-float64(int(x)) > 0 || y-float64(int(y)) > 0 {
		return 0, false
	}

	return x*3 + y, true
}

func readInput(input string) []config {
	blocks := strings.Split(input, "\n\n")
	configs := make([]config, 0, len(blocks))

	for _, b := range blocks {
		if b == "" {
			continue
		}

		var (
			cfg   config
			lines = strings.Split(b, "\n")
		)
		fmt.Sscanf(lines[0], "Button A: X+%f, Y+%f", &cfg.ButtonA.X, &cfg.ButtonA.Y)
		fmt.Sscanf(lines[1], "Button B: X+%f, Y+%f", &cfg.ButtonB.X, &cfg.ButtonB.Y)
		fmt.Sscanf(lines[2], "Prize: X=%f, Y=%f", &cfg.Prize.X, &cfg.Prize.Y)

		configs = append(configs, cfg)
	}

	return configs
}
