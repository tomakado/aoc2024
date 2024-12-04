package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

var (
	opPtn     = regexp.MustCompile(`(don't|do|mul\((?P<a>[0-9]{1,3}),(?P<b>[0-9]{1,3})\))`)
	opArgsPtn = regexp.MustCompile(`(?P<a>[0-9]{1,3}),(?P<b>[0-9]{1,3})`)
)

func main() {
	var (
		sum          int
		isMulEnabled = true
		matches      = opPtn.FindAll(input, -1)
	)

	for _, m := range matches {
		mStr := string(m)

		switch {
		case strings.HasPrefix(mStr, "don't"):
			isMulEnabled = false
			continue
		case strings.HasPrefix(mStr, "do"):
			isMulEnabled = true
			continue
		case !isMulEnabled:
			continue
		}

		argsTogether := opArgsPtn.Find(m)
		aStr, bStr, _ := strings.Cut(string(argsTogether), ",")
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)

		sum += a * b
	}

	fmt.Println(sum)
}
