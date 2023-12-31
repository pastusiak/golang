package console

import (
	"os"
)

func WithArgs() (isArgs bool, option []string, quit bool) {
	isArgs = false
	quit = false

	args := os.Args
	if len(args) <= 1 {
		return
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--help", "-?":
			option = append(option, "?")
			isArgs = true
			quit = true
		case "-q":
			isArgs = true
			quit = true
		case "-o":
			if len(args) >= (i + 1) {
				option = append(option, args[(i+1)])
				i++
				isArgs = true
			}
		}
	}

	return
}
