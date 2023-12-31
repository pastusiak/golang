package console

import (
	"strings"
)

func ClearInput(input string, toLower bool) string {
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)

	if toLower {
		input = strings.ToLower(input)
	}

	return input
}
