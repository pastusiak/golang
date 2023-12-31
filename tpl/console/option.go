
package console

import (
	"bufio"
	"fmt"
	"os"
)

func SelectOption(option string, cmdList []CmdInterface) (result bool, shutdown bool) {
	result = false
	shutdown = false

	option = ClearInput(option, true)

	if option == "?" {
		for _, item := range cmdList {
			name, desc := item.CmdConfig()
			fmt.Printf("| [%s] %s\n\r", name, desc)
		}
		result = true
		shutdown = false
		return
	}

	if option == "q" {
		fmt.Print("Are you sure? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		prompt, _ := reader.ReadString('\n')
		prompt = ClearInput(prompt, true)
		if prompt == "y" {
			result = true
			shutdown = true
		} else {
			result = false
			shutdown = false
		}
		return
	}

	for _, item := range cmdList {
		name, _ := item.CmdConfig()
		if name == option {
			result = item.CmdExecute()
			shutdown = false
			return
		}
	}

	return
}
