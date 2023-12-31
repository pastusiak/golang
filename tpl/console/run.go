package console

import (
	"bufio"
	"log"
	"os"
)

var cmdList []CmdInterface

func RegisterCmd(cmd CmdInterface) {
	cmdList = append(cmdList, cmd)
}

func Run() bool {
	log.Println("Application start")

	isArgs, options, quit := WithArgs()

	if isArgs {
		log.Printf("Run with params. Use option: %s\n\r", options)

		quitAfterRun := false

		for i := 0; i < len(options); i++ {
			_, quitByOption := SelectOption(options[i], cmdList)

			if quitByOption {
				quitAfterRun = true
			}
		}

		if quit || quitAfterRun {
			log.Println("Application shutdown")
			return true
		}
	}

	for {
		ShowMenu()

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		_, quit := SelectOption(option, cmdList)

		if quit {
			break
		}
	}

	log.Println("Application shutdown")
	return true
}
